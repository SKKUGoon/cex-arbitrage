from cex.cex_factory import CexFactoryX, CexManagerX
from .sig1_gen import gen_signal_notice
from utility.coloring import PrettyColors


class StrategyIEXANotice:
    """
    Inter Exchange Arbitrage Strategy
    In this case,
      exchange_long: upbit
      exchange_short: binance
    
    Additional strategy. If upbit blocks some currency,
    premium on that coin will sore.
    """
    def __init__(self, exchange_long: CexManagerX, exchange_short: CexManagerX) -> None:
        self.exchange = CexFactoryX()
        self.longX = exchange_long
        self.shortX = exchange_short

    def target_assets(self, long_key_currency: str, short_key_currency: str) -> set:
        PrettyColors().print_ok_cyan(val="Update common traded asset")
        l = self.exchange.get_tradable(self.longX)[long_key_currency.upper()]
        s = self.exchange.get_tradable(self.shortX)[short_key_currency.upper()]

        l = set(l)
        result = set()
        for s_asset in s:
            if s_asset in l:
                result.add(s_asset)
        return result
    
    @staticmethod
    def run(env: str, hostname: str, common_trade: set):
        gen_signal_notice(env, hostname, common_trade)