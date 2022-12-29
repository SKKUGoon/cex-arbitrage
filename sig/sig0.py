from cex.cex_factory import CexFactoryX, CexManagerX
from .sig0_gen import gen_signal_iexa_multi
from utility.coloring import PrettyColors

import multiprocessing as mp
from typing import Set, Callable
import sys


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
        PrettyColors().print_ok_cyan("Update common traded asset")
        l = self.exchange.get_tradable(self.longX)[long_key_currency.upper()]
        s = self.exchange.get_tradable(self.shortX)[short_key_currency.upper()]
        
        l = set(l)
        result = set()  # Long exchange(i.e. Upbit) only
        for s_asset in s:
            if s_asset in l:
                result.add(s_asset)
        return result
    
    def run_multi(self, x_long: dict, x_short: dict, assets: set, 
            ws_func_long: Callable, ws_func_short: Callable, 
            ws_keycurr_long: str, ws_keycurr_short: str,
            env: str="dev", hostname: str="localhost"):
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
        PrettyColors().print_ok_cyan(f"Created {len(assets)}# queues")
        
        # Set up websocket workers
        p1 = mp.Process(
          target=ws_func_long, 
          args=(x_long, multiq_long, ws_keycurr_long,))
        p2 = mp.Process(
          target=ws_func_short, 
          args=(x_short, multiq_short, ws_keycurr_short,)
        )

        # Set up as daemon process
        p1.daemon = True
        p2.daemon = True
        PrettyColors().print_ok_cyan(f"Created p1, p2. Both as Daemon")

        p1.start()
        p2.start()

        gen_signal_iexa_multi(assets, multiq_long, multiq_short, hostname, env)
        PrettyColors().print_ok_green("System exit. Daemon Process terminated")
        sys.exit()
        return
