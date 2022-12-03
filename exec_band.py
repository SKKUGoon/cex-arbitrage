from cex.cex_factory import CexFactoryX
from cex.binance_future import BinanceFutureX
from cex.domestic import UpbitX
from sig.sig0 import StrategyIEXA
from sig.sig0_gen import gen_band_iexa

import time
import argparse


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    # -env ENVIRONMENT -host HOSTNAME
    parser.add_argument("-env", "--environment", help="Type of Environment, Dev or Deploy")
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")

    args = parser.parse_args()

    cm = CexFactoryX()
    binance_x = BinanceFutureX()
    upbit_x = UpbitX()
    cm.get_connection(binance_x)
    cm.get_connection(upbit_x)

    strat = StrategyIEXA(upbit_x, binance_x)
    a = strat.target_assets('krw', 'usdt')
    
    gen_band_iexa(
        a,
        exchange_long=upbit_x,
        exchange_short=binance_x,
        env=args.environment,
        hostname=args.hostname
    )

    