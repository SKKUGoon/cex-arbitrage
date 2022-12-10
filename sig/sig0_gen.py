from cex.cex_factory import CexFactoryX, CexManagerX
from utility.fx import forex
from utility.bollinger import bollinger
from utility.coloring import PrettyColors

from typing import Final
import multiprocessing as mp
import requests
import time


def gen_signal_iexa_multi(assets: set, qs_long: dict, qs_short: dict, 
        env: str="dev", hostname: str="localhost", data_collect: int=30):
    """
    @param qs_long, qs_short: they look like as such. 
    { <asset name>: multiprocessing.Queue, ... }
    @param data_collect: how much time period between each premium calc burst.(secs)
    """
    print(PrettyColors.WARNING + "GEN_SIGNAL_IEXA_MULTI infinite loop start" + PrettyColors.ENDC)
    # Signal burst in seconds
    if env.lower() == "dev":
        SIGNAL_TO = f"http://{hostname}:10532/premium"
    elif env.lower() == "deploy":
        SIGNAL_TO = f"http://{hostname}:10532/premium"
    else:
        raise RuntimeError(f"Environment `{env}` is not one of the specified.")

    start = time.time()
    # Websocket packet handling endless cycle.
    # Stops after explicit ctrl + c key input
    while True:
        # Data collecting time ticker.
        # It calls for `forex()` value every `data_collect` seconds
        #   `forex()` will return foreign exchange value. 
        #   It is assumed that 1 USD will have equal amount of value to 1 USDT
        collect = time.time() - start >= data_collect
        if collect:
            fx_value, true_val = forex()
            if not true_val:
                print(PrettyColors.WARNING + "Foreign Exchange Rate, Approximated", PrettyColors.ENDC)

        # From the set it process it waits for the websocket data
        #   to be inserted in a queue. It waits for `timeout` amount
        #   of 2 seconds. Same goes for both exchanges. With the information
        #   of both exchanges, calculate premium.(if statement) If one of the 
        #   queue is empty, the data cannot be calculated. (else statement)
        packet = {
            "type": "iexa",
            "data": {
                "exchange_pair": {
                    "long": "upbit",
                    "short": "binance"
                },
                "asset_premium": list()
            }
        }
        for a in assets:
            # Data extraction from queues
            try:
                l = qs_long[a].get(timeout=2)
            except Exception as e:
                print(PrettyColors.WARNING + f"Asset {a}:: data queue LONG empty" + PrettyColors.ENDC)
                l = None
            try :
                s = qs_short[a].get(timeout=2)
            except Exception as e:
                print(PrettyColors.WARNING + f"Asset {a}:: data queue SHORT empty" + PrettyColors.ENDC)
                s = None

            if collect:
                # Why `if collect` after `try - except` on data extraction from queues
                # We have to extract old information and use the information on time. 
                start = time.time()
                if l is not None and s is not None:
                    # Premium is calculated as such
                    #   Denote 1) Best Asking Price: BAP, 2) Best Biding price BBP
                    # Calculations:
                    #   (FX adjusted BAP / BBP) - 1
                    #   If the calculation of premium is lower than certain threshold, 
                    #     the backend server will tell the traders to submit orders.
                    # This function will only <b>convey</b> the premium to the server.
                    # The backend server will receive the premium information by `packet`
                    #   `packet` is a dictionary with "type", "data"-"exchangepair", 
                    #   "data"-"asset_premium" as a key. 
                    premium = (
                        (l['orderbook_units'][0]['ask_price'] / fx_value) - float(s['b'])
                    ) / float(s['b'])
                    packet["data"]["asset_premium"].append(
                        {
                            "asset": a,
                            "premium": premium
                        }
                    )
                    print(PrettyColors.OKGREEN + f"Premium {a} added" + PrettyColors.ENDC)
                else:
                    print(PrettyColors.FAIL + f"Premium {a}: No Calc" + PrettyColors.ENDC)
        
        if collect:
            print(PrettyColors.BOLD + packet.__repr__() + PrettyColors.ENDC)
            resp = requests.post(SIGNAL_TO, json=packet)
            if resp.status_code == 200:
                print(
                    PrettyColors.OKGREEN 
                    + f"Status Code{resp.status_code}: {resp.json()}"
                    + PrettyColors.ENDC
                )
            else:
                print(
                    PrettyColors.FAIL 
                    + f"Status Code{resp.status_code}: {resp.json()}"
                    + PrettyColors.ENDC
                )
            

def gen_band_iexa(tickers: set, exchange_long: CexManagerX, exchange_short: CexManagerX, env: str="dev", hostname: str="localhost"):
    """
    @param exchange_long, exchange_short: Class type CexManagerX 
    @param burst_interval_min: Signal bursting request by minutes
    """
    print(PrettyColors.WARNING + "GEN_BAND_IEXA start" + PrettyColors.ENDC)
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
                + PrettyColors.ENDC
            )
        else:
            print(
                PrettyColors.FAIL 
                + f"{resp.status_code} {resp.json()['data']['message']}"
                + PrettyColors.ENDC
            )
        print(post_premium_data)
    # End of update
    print(PrettyColors.WARNING + "GEN_BAND_IEXA end" + PrettyColors.ENDC)
        
        