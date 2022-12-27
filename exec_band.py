from cex.cex_factory import CexFactoryX
from cex.binance_future import BinanceFutureX
from cex.domestic import UpbitX
from utility.coloring import PrettyColors
from sig.sig0 import StrategyIEXA
from sig.sig0_gen import gen_band_iexa

import time
import argparse


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    # -env ENVIRONMENT -host HOSTNAME
    parser.add_argument("-env", "--environment", help="Type of Environment, Dev or Deploy")
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")
    parser.add_argument("-upbitkey", "--upbitkeycurrency", help="Upbit's key currency")
    parser.add_argument("-binancekey", "--binancekeycurrency", help="Binance's key currency")

    args = parser.parse_args()
    
    INTERVAL = 60 * 5  # secs
    job_start = time.time()

    cm = CexFactoryX()
    binance_x = BinanceFutureX()
    upbit_x = UpbitX()

    binance_key_currency = args.binancekeycurrency  # USDT or BUSD
    upbit_key_currency = args.upbitkeycurrency  # KRW


    cm.get_connection(binance_x)
    cm.get_connection(upbit_x)

    strat = StrategyIEXA(upbit_x, binance_x)
    a = strat.target_assets(upbit_key_currency, binance_key_currency)
    
    gen_band_iexa(
        a,
        exchange_long=upbit_x,
        exchange_short=binance_x,
        env=args.environment,
        hostname=args.hostname
    )
    job_done = time.time()
    PrettyColors().print_warning(
        val=f"Band Updated. {time.strftime('%c', time.localtime())}"
    )
    # Wait for 5 minute. After 5 minute, container restart.
    time.sleep(60 * 5 - (job_done - job_start))  # 60sec * 5 = 5 minutes
