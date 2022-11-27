from cex.cex_factory import CexFactoryX
from cex.binance import BinanceX
from cex.domestic import UpbitX
from sig.sig0 import StrategyIEXA
from utility.graceful_shutdown import handle_ctrlc

import signal


if __name__ == "__main__":
    cm = CexFactoryX()
    binance_x = BinanceX()
    upbit_x = UpbitX()

    strat = StrategyIEXA(upbit_x, binance_x)
    a = strat.target_assets('krw', 'usdt')

    u_a = cm.gen_ticker(upbit_x, a, "krw")
    b_a = cm.gen_ticker(binance_x, a, "usdt")

    print(u_a)
    print(b_a)

    signal.signal(signal.SIGINT, handle_ctrlc)
    strat.run()
