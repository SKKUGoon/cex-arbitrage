from cex.cex_factory import CexFactoryX
from cex.binance_future import BinanceFutureX
from cex.domestic import UpbitX
from cex.ws_binance import binance_ws_multi
from cex.ws_upbit import upbit_ws_multi
from sig.sig0 import StrategyIEXA
from utility.graceful_shutdown import handle_ctrlc

import signal
import argparse


if __name__ == "__main__":
    # Parse flags for execution
    parser = argparse.ArgumentParser()
    # -env ENVIRONMENT -host HOSTNAME
    parser.add_argument("-env", "--environment", help="Type of Environment, Dev or Deploy")
    parser.add_argument("-host", "--hostname", help="Name of host. Service name such as trade_control")

    args = parser.parse_args()

    cm = CexFactoryX()
    binance_x = BinanceFutureX()
    upbit_x = UpbitX()

    strat = StrategyIEXA(upbit_x, binance_x)
    a = strat.target_assets('krw', 'usdt')

    u_a = cm.gen_ticker(upbit_x, a, 'krw')
    b_a = cm.gen_ticker(binance_x, a, 'usdt')

    signal.signal(signal.SIGINT, handle_ctrlc)
    strat.run_multi(
        x_long=u_a,
        x_short=b_a,
        assets=a,
        ws_func_long=upbit_ws_multi,
        ws_func_short=binance_ws_multi,
        env=args.environment,
        hostname=args.hostname
    )
