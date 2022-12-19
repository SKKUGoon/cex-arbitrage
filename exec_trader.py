from trader.order_process import ArbitrageIEXA
from cex.binance_future_trade import BinanceFutureT
from cex.domestic_trade import UpbitT

import argparse


if __name__ == "__main__":
    # Parse flags for execution
    parser = argparse.ArgumentParser()
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")
    parser.add_argument("-upbitkey", "--upbitkeycurrency", help="Upbit's key currency")
    parser.add_argument("-binancekey", "--binancekeycurrency", help="Binance's key currency")
    parser.add_argument("-upbitir", "--upbitinvestratio", help="Upbit's invest ratio wrt balance")
    parser.add_argument("-binanceir", "--binanceinvestratio", help="Binance's invest raio wrt balance")

    args = parser.parse_args()

    binance_t = BinanceFutureT(
        args.binancekeycurrency.upper()
    )
    upbit_t = UpbitT(
        args.upbitkeycurrency.upper()
    )

    upbit_ir = float(args.upbitinvestratio)
    binance_ir = float(args.binanceinvestratio)

    arbitrager = ArbitrageIEXA(
        exchange_long=upbit_t, 
        exchange_short=binance_t, 
        long_total_use=upbit_ir, 
        short_total_use=binance_ir,
        hostname=args.hostname
    )
    arbitrager.listen("trade_channel", None, None)
  
