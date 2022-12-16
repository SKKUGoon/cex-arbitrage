from trader.order_process import ArbitrageIEXA
from cex.binance_future_trade import BinanceFutureT
from cex.domestic_trade import UpbitT

import argparse


if __name__ == "__main__":
    # Parse flags for execution
    parser = argparse.ArgumentParser()
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")

    args = parser.parse_args()

    binance_t = BinanceFutureT()
    upbit_t = UpbitT()

    arb = ArbitrageIEXA(binance_t, upbit_t, args.hostname)
    arb.listen("trade_channel")