from cex.cex_factory_trade import CexManagerT
from .order_report import report_order
from utility.hedge import hedge
from utility.coloring import PrettyColors


def iexa_enter_pos(mq_data: dict, lev: int, balance: dict, order_ratio: float, long_ex: CexManagerT, short_ex: CexManagerT):
    if iexa_check_pos(mq_data, long_ex, short_ex):
        print(
            PrettyColors.WARNING
            + "Signaled asset already present, No order sent"
            + PrettyColors.ENDC
        )
        # If asset position is already present, discontinue
        return

    # Both market enter-position-order for an asset
    asset = mq_data['a']
    # Calculate with hedge
    lo_money = balance['l'] * order_ratio
    lo_quantity = lo_money / mq_data["pl"]
    so_quantity = hedge(lo_quantity, lev)
    # so_money = balance['s'] * order_ratio
    report_order("enter", asset, lo_quantity, so_quantity)

    long_ex.conn.create_order(
        f"{asset}/{long_ex.EX_CURRENCY}",
        "market",
        "buy",
        lo_quantity,
        mq_data["pl"]
    )
    # Adjust leverage for IEXA arb
    short_ex.conn.fapiPrivate_post_leverage({
        "symbol": f"{asset}{short_ex.EX_CURRENCY}".upper(),
        "leverage": int(lev),
    })
    short_ex.conn.create_order(
        f"{asset}/{short_ex.EX_CURRENCY}",
        "market",
        "sell",
        so_quantity,
        mq_data["ps"],
    )

def iexa_check_pos(mq_data: dict, long_ex: CexManagerT, short_ex: CexManagerT):
    """
    Each `CexManagerT` object has method `balance`. `balance` returns 
    "open_position_set" keys.

    return True: If asset in both long_exchange and short_exchange balance
    return False: Else
    """
    asset = mq_data["a"]

    lb = long_ex.balance(long_ex.EX_CURRENCY)
    sb = short_ex.balance(short_ex.EX_CURRENCY)
    lb_has_asset = asset in lb['open_position_set']
    sb_has_asset = asset in sb['open_position_set']
    if lb_has_asset and sb_has_asset:
        return True
    else:
        return False

def iexa_exit_pos(mq_data: dict, long_ex: CexManagerT, short_ex: CexManagerT):
    if not iexa_check_pos(mq_data, long_ex, short_ex):
        # If there's no position to exit, discontinue
        return
    
    # Both market exit-position-order for an asset
    asset = mq_data["a"]
    
    long_ex.conn.create_order(
        f"{asset}/{long_ex.EX_CURRENCY}",
        "market",
        "sell",
        upbit_pos_balance(long_ex, f"{asset}".upper()),
        mq_data["pl"]
    )
    short_ex.conn.create_market_buy_order(
        symbol=f"{asset}/{short_ex.EX_CURRENCY}",
        amount=binance_pos_balance(short_ex, f"{asset}{short_ex.EX_CURRENCY}"),
        params={"reduceOnly": True}
    )

def binance_pos_balance(ex: CexManagerT, tgt: str) -> str:
    bals = ex.conn.fetch_balance()
    bals = bals["info"]["positions"]
    for i in bals:
        if i["symbol"] == tgt:
            return i["positionAmt"]

def upbit_pos_balance(ex: CexManagerT, tgt: str) -> str:
    bals = ex.conn.fetch_balance()
    bals = bals["info"]
    for i in bals:
        if i["currency"] == tgt:
            return i["balance"]
