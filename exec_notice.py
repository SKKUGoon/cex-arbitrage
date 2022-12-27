from cex.binance_future import BinanceFutureX
from cex.domestic import UpbitX
from sig.sig1 import StrategyIEXANotice
from utility.coloring import PrettyColors

import argparse
import time


if __name__ == "__main__":
    # Parse flags for execution
    parser = argparse.ArgumentParser()
    # -env ENVIRONMENT -host HOSTNAME
    parser.add_argument("-env", "--environment", help="Type of Environment, Dev or Deploy")
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")
    parser.add_argument("-upbitkey", "--upbitkeycurrency", help="Upbit's key currency")
    parser.add_argument("-binancekey", "--binancekeycurrency", help="Binance's key currency")

    args = parser.parse_args()

    job_start = time.time()

    binance_x = BinanceFutureX()
    upbit_x = UpbitX()

    binance_key_currency = args.binancekeycurrency  # USDT or BUSD
    upbit_key_currency = args.upbitkeycurrency  # KRW

    strat = StrategyIEXANotice(upbit_x, binance_x)
    a = strat.target_assets(upbit_key_currency, binance_key_currency)
    print(a)

    strat.run(
        env=args.environment,
        hostname=args.hostname,
        common_trade=a
    )

    job_done = time.time()
    PrettyColors().print_warning(
        val=f"Notice Updated. {time.strftime('%c', time.localtime())}\n"
    )
    # Wait for 1 minute. After 1 minute, container restart.
    time.sleep(60 * 1 - (job_done - job_start))