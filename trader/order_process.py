from cex.cex_factory_trade import CexManagerT, CexFactoryT
from .pair_pos import iexa_enter_pos, iexa_exit_pos
from utility.hedge import Leverage
from utility.parse_yaml import ConfigParse
from utility.coloring import PrettyColors
from utility.fx import forex

import redis

import warnings
import json

class ArbitrageIEXA:
    """
    Inter Exchange Arbitrage Strategy
    In this case,
      exchange_long: upbit - KRW
      exchange_short: binance - USDT or BUSD
        - It can use leverage

    Whole container will be restarted after 24 hours. 
    At the start of the trader object class, 
    """
    def __init__(self, exchange_long: CexManagerT, exchange_short: CexManagerT, 
        long_total_use: float, short_total_use: float, hostname: str):
        # Exchange manager
        self.exchange = CexFactoryT()
        self.long = exchange_long
        self.short = exchange_short

        # Trade information
        self.multi_bal = self._get_balance()
        self.day_leverage = self.set_leverage(long_total_use, short_total_use)

        r = redis.Redis(host=hostname, port=6379, db=0, password="mypassword")
        self.pubsub = r.pubsub()

    def _get_balance(self):
        """
        Return balance (forex adjusted). 1) Get balance by using keyword 
        `key_balance` > `free`, 2) Get foreign exchange rate. 
        """
        fx, _ = forex()
        lb = self.exchange.get_balance(self.long, self.long.EX_CURRENCY)
        sb = self.exchange.get_balance(self.short, self.short.EX_CURRENCY)
        return {
            "l": lb["key_balance"]["balance"]["free"],
            "s": sb["key_balance"]["balance"]["free"],
            "fx": fx
        }

    def set_leverage(self, ex_long_usage: float, ex_short_usage: float) -> int:
        """
        Read balance from `self.multi_bal` and return leverage rate in integer. 
        If leverage is 1) not within bounds of test.py it issues a warning. 
        If leverage is LEQ than 0, it issues a RuntimeError. 

        `STANDARD` wallet is "wallet1". In this case `STANDARD` wallet is Upbit.
        (Because Upbit doesn't offer any leverage service.)
        """
        # Get PRESET max-min leverage
        cp = ConfigParse("./exchange.yaml")
        d = cp.parse()
        lev_min, lev_max = d["rule"]["minimum_leverage"], d["rule"]["maximum_leverage"]
        
        # Calculate the day's leverage
        STANDARD, COMPARISON = "wallet1", "wallet2"
        lev = Leverage(
            wallet1_balance=self.multi_bal['l'] / self.multi_bal['fx'],
            wallet2_balance=self.multi_bal['s'],
        )
        app_lev = lev.rcmd_leverage(STANDARD, ex_long_usage, ex_short_usage)
        if app_lev[COMPARISON] <= 0:
            raise RuntimeError(f"leverage should be greater than 0. {app_lev[COMPARISON]} <= 0")
        try:
            assert lev_min <= app_lev[COMPARISON] <= lev_max
        except AssertionError:
            warnings.warn(f"Warning. Leverage out of bound {app_lev[COMPARISON]}~[{lev_min}, {lev_max}]")
        
        return app_lev[COMPARISON]

    def listen(self, channel_name: str):
        """
        Infinte loop that has 2 callback functions
          - `callback_enter_pos`
          - `callback_exit_pos`
        Subscribe + listen to redis pubsub message queue. Translate b'' to json. 
        If order type is (long ex-buy & short ex-sell), just buy at market price. This is
        executed by calling `callback_enter_pos` function.  
        If order type is (long ex-sell & short ex-buy), first !CHECK IF THE POSITION 
        IS PRESENT!, and then execute the order by calling `callback_exit_pos` function. 
        """
        self.pubsub.subscribe(channel_name)

        while True:            
            res = self.pubsub.get_message(timeout=5)
            if res is not None:
                if res['type'] == "subscribe":
                    print(
                        PrettyColors.OKGREEN
                        + f"Subcribed to channel: {res['channel'].decode('utf-8')}"
                        + PrettyColors.ENDC
                    )
                elif res['type'] == "message":    
                    try:
                        data = res['data'].decode('utf-8')
                    except Exception as e:
                        print(PrettyColors.FAIL + e + PrettyColors.ENDC)
                        continue

                    try:
                        jdata = json.loads(data)
                    except Exception as e:
                        print(PrettyColors.FAIL + e + PrettyColors.ENDC)
                        continue

                    if jdata["t"] == "enter":
                        is_exec = iexa_enter_pos(
                            mq_data=jdata,
                            lev=self.day_leverage,
                            balance=self.multi_bal,
                            order_ratio=0.1,
                            long_ex=self.long,
                            short_ex=self.short,
                        )
                        if not is_exec:
                            # Stop calling binance asset when no trade is made
                            PrettyColors().print_ok_blue(val="No trade made")
                            continue
                        self.multi_bal = self._get_balance()
                        
                    elif jdata["t"] == "exit":
                        is_exec = iexa_exit_pos(
                            mq_data=jdata,
                            long_ex=self.long,
                            short_ex=self.short,
                        )
                        if not is_exec:
                            # Stop calling binance asset when no trade is made
                            PrettyColors().print_ok_blue(val="No trade made")
                            continue
                        self.multi_bal = self._get_balance()
                        
                    else:
                        print(PrettyColors.FAIL + "wrong message" + PrettyColors.ENDC)


        