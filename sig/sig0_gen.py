from cex.cex_factory import CexFactoryX, CexManagerX
from utility.fx import forex
from utility.bollinger import bollinger
from utility.coloring import PrettyColors

from typing import Final
import time
import json

import requests
import redis


def gen_signal_iexa_multi(assets: set, qs_long: dict, qs_short: dict, 
        hostname: str, env: str="dev", data_collect: int=30):
    """
    Infinite Loop. To be closed with `ctrl+c`. Data collecting after time ticker.

    1) About foreign exchange rate
        It calls for `forex()` value every `data_collect` seconds. 
        `forex()` will return foreign exchange value. It is assumed that 1 USD will have equal amount of value to 1 USDT.
    2) About Websocket data processing
        It is to be inserted in a queue - waits for `timeout` seconds (default: 2).
        Same goes for both exchanges. With the information of both exchanges, calculate
        premium. If one of the queue is empty, the data cannot be calculated.

        2-1) Try-except statement exist because we have to extract data from the queue
        after timeout. 
    3) About premium calculation
        BAP: Best Asking price
        BBP: Best Bidding price
        premium = {(FX adjusted BAP) / BBP} - 1
        All premium will be sent via redis pubsub as a json data drop packet. 

        3-1) If premium is lower than certain threshold, the backend server will tell the 
        trader channel to submit orders. The backend server will receive the premium 
        information by `packet`. `packet` is a dictionary with `type`, `data`-
        `exchangepair`, `data`-`asset_premium` as a key. 


    @param qs_long, qs_short: they look like as such. 
    { <asset name>: multiprocessing.Queue, ... }
    @param data_collect: how much time period between each premium calc burst.(secs)
    """
    print(
        PrettyColors.WARNING 
        + "GEN_SIGNAL_IEXA_MULTI infinite loop start" 
        + PrettyColors.ENDC,
        flush=True
    )
    # Signal burst in seconds
    if env.lower() == "dev":
        r = redis.Redis(host=hostname, port=6379, db=0)
    elif env.lower() == "deploy":
        r = redis.Redis(host=hostname, port=6379, db=0, password="mypassword")
    else:
        raise RuntimeError(f"environment env={env} is not one of the specified")

    start = time.time()
    while True:
        collect = time.time() - start >= data_collect
        if collect:
            # Handle 1)
            fx_value, true_val = forex()
            if not true_val:
                print(
                    PrettyColors.WARNING + 
                        "Foreign Exchange Rate, Approximated" + 
                        PrettyColors.ENDC,
                    flush=True
                )

        packet = {
            "type": "iexa",
            "status": True,
            "data": {
                "exchange_pair": {
                    "long": "upbit",
                    "short": "binance"
                },
                "asset_premium": dict()
            }
        }
        # Handle 2)
        for a in assets:
            # Handle 2-1)
            try:
                l = qs_long[a].get(timeout=2)
            except Exception as e:
                print(
                    PrettyColors.WARNING 
                    + f"Asset {a}:: data queue LONG empty" 
                    + PrettyColors.ENDC,
                    flush=True
                )
                l = None
            try :
                s = qs_short[a].get(timeout=2)
            except Exception as e:
                print(
                    PrettyColors.WARNING 
                    + f"Asset {a}:: data queue SHORT empty" 
                    + PrettyColors.ENDC,
                    flush=True
                )
                s = None
            if not collect:
                continue

            if l is not None and s is not None:
                # Handle 3)
                premium = (
                    (l['orderbook_units'][0]['ask_price'] / fx_value) - float(s['b'])
                ) / float(s['b'])
                packet["data"]["asset_premium"] = {
                    "asset": a,
                    "premium": premium,
                }
                # Handle 3-1)
                packet_dump = json.dumps(packet)
                r.publish(channel="signal_channel", message=packet_dump)
                print(
                    PrettyColors.OKGREEN 
                    + f"Premium {a} published" 
                    + PrettyColors.ENDC,
                    flush=True
                )
            else:
                print(
                    PrettyColors.FAIL 
                    + f"Premium {a}: No Calc" 
                    + PrettyColors.ENDC,
                    flush=True
                )
            # Reset time.
            start = time.time()
            

def gen_band_iexa(tickers: set, exchange_long: CexManagerX, exchange_short: CexManagerX, env: str="dev", hostname: str="localhost"):
    """
    @param exchange_long, exchange_short: Class type CexManagerX 
    @param burst_interval_min: Signal bursting request by minutes
    """
    print(
        PrettyColors.WARNING + "GEN_BAND_IEXA start" + PrettyColors.ENDC,
        flush=True
    )
    cm = CexFactoryX()
    # Signal burst in seconds
    if env.lower() == "dev":
        BURST_TO = f"http://{hostname}:10532/band"
    elif env.lower() == "deploy":
        BURST_TO = f"http://{hostname}:10532/band"
    else:
        raise RuntimeError(f"Environment `{env}` is not one of the specified.")
    SPLIT_BY: Final = 5
    
    # Split tickers `tickers` by length of `5`
    ticker_ls = list(tickers)
    split_t = list()
    for i in range(0, len(ticker_ls), SPLIT_BY):
        split_t.append(set(ticker_ls[i: (i+SPLIT_BY)]))

    # Get history and fx information

    fx, _ = forex()
    for sub_t in split_t:
        # Execute - 1) Calculate premium and 2) Request for each batch `sub_t`
        post_premium_data = list()
        for t in sub_t:       
            # Get closing data
            l = cm.get_history(exchange_long, t, "KRW")  # KRW since its upbit
            s = cm.get_history(exchange_short, t, "USDT")  # USDT since its

            # Calculate premium
            premiums = list()
            for l, s, in zip(l, s):
                premiums.append(((l / fx) - s) / s)
            upper, lower = bollinger(premiums)

            # Premium upload for each batch's element `t`
            post_premium_data.append(
                {
                    "asset": t,
                    "upper": upper,
                    "lower": lower,
                }
            )
        
        # Premium upload to database
        resp = requests.post(BURST_TO, json={
            "type": "iexa",
            "data": post_premium_data
        })
        if resp.status_code == 200:
            print(
                PrettyColors.OKGREEN 
                + f"{resp.json()['type']} {resp.json()['data']['message']}"
                + PrettyColors.ENDC,
                flush=True
            )
        else:
            print(
                PrettyColors.FAIL 
                + f"{resp.status_code} {resp.json()['data']['message']}"
                + PrettyColors.ENDC,
                flush=True
            )
        print(post_premium_data)
    # End of update
    print(
        PrettyColors.WARNING 
        + "GEN_BAND_IEXA end" 
        + PrettyColors.ENDC,
        flush=True
    )
        
        