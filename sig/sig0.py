from cex.cex_factory import CexFactoryX, CexManagerX
from cex.ws_binance import binance_ws
from cex.ws_upbit import upbit_ws
from .sig0_gen import gen_signal_iexa, gen_signal_iexa_multi
from utility.coloring import PrettyColors

import multiprocessing as mp
from typing import Set, Callable


class StrategyIEXA:
    """
    Inter Exchange Arbitrage Strategy
    In this case,
      exchange_long: upbit
      exchange_short: binance
    """
    def __init__(self, exchange_long: CexManagerX, exchange_short: CexManagerX) -> None:
        self.exchange = CexFactoryX()
        self.longX = exchange_long
        self.shortX = exchange_short

    def target_assets(self, long_key_currency: str, short_key_currency: str) -> Set:
        """
        @param long_key_currency: Key currency that's traded in 
          Long-position only exchange
        @param short_key_currency: Key currency that's traded in 
          Short-position only exchange

        Find common assets pairs between exchange_long and exchange_short.
        Returns common currency in set
        """
        print(PrettyColors.OKBLUE + "Update common traded asset" + PrettyColors.ENDC)
        l = self.exchange.get_tradable(self.longX)[long_key_currency.upper()]
        s = self.exchange.get_tradable(self.shortX)[short_key_currency.upper()]
        
        l = set(l)
        result = set()  # Long exchange(i.e. Upbit) only
        for s_asset in s:
            if s_asset in l:
                result.add(s_asset)
        return result
    
    def run_multi(self, x_long: dict, x_short: dict, assets: set, 
            ws_func_long: Callable, ws_func_short: Callable):
        """
        @param x_long: dictionary that contains websocket subscription message 
          for Long-position only exchange (upbit)
        @param x_short: dictionary that contains websocket subscription message
          for Short-position only exchange (binance)
        @param assets: hash-set that contains common tickers between
          Long-position exchange and Short-position exchange
        @param ws_func_long: Long-position exchange (upbit) websocket handler
        @param ws_func_short: Short-position exchange (binance) websocket handler

        This function will run forever until the ctrl+c is called.
        Function will generate 2 dictionaries that contain multiprocessing.Queue * k, 
        where k is the length of assets - one for long-exchange and the other for short.

        Must be run in __main__.
        """
        multiq_long, multiq_short = dict(), dict()
        for asset in assets:
            multiq_long[asset]: mp.Queue = mp.Queue()
            multiq_short[asset]: mp.Queue = mp.Queue()
        print(PrettyColors.OKBLUE + f"Created {len(assets)}# queues" + PrettyColors.ENDC)
        
        p1 = mp.Process(target=ws_func_long, args=(x_long, multiq_long,))
        p2 = mp.Process(target=ws_func_short, args=(x_short, multiq_short,))
        p3 = mp.Process(target=gen_signal_iexa_multi, args=(assets, multiq_long, multiq_short,))

        p1.start()
        p2.start()
        p3.start()

        p1.join()
        p2.join()
        p3.join()  # Wait for completion

    def run(self):
        """
        There is an API request restriction in Binance API.
        Since there are so many coin, we are going to use websocket.
        Subscribe to currencies in asset_ls. This function will 
        run forever until the ctrl+c is called. 

        Subscribe to both exchange and it will store information 
        inside a queue. The information will be processed each minute.

        Must be run in __main__.
        """
        q_upbit: mp.Queue = mp.Queue()
        q_binance: mp.Queue = mp.Queue()

        p1 = mp.Process(target=upbit_ws, args=(q_upbit,))
        p2 = mp.Process(target=binance_ws, args=(q_binance,))
        p3 = mp.Process(target=gen_signal_iexa, args=(q_upbit, q_binance,))

        p1.start()
        p2.start()
        p3.start()

        p1.join()
        p2.join()
        p3.join()  # Wait for completion
