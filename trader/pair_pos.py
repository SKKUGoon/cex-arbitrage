from cex.cex_factory_trade import CexManagerT
from .order_report import report_order
from utility.hedge import hedge
from utility.coloring import PrettyColors


def iexa_enter_pos(mq_data: dict, lev: int, balance: dict, order_ratio: float, long_ex: CexManagerT, short_ex: CexManagerT) -> bool:
    if iexa_check_pos(mq_data, long_ex, short_ex):
        # If asset position is already present, discontinue
        PrettyColors().print_warning("Signaled asset ALREADY present, No orders sent")
        return False

    # Both market enter-position-order for an asset
    asset = mq_data['a']
    # Calculate with hedge
    lo_money = balance['l'] * order_ratio
    lo_quantity = lo_money / mq_data["pl"]
    so_quantity = hedge(lo_quantity, lev)
    p = mq_data["pm"]
    # so_money = balance['s'] * order_ratio
    report_order("enter", asset, lo_quantity, so_quantity, p)

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
    return True

def iexa_check_pos(mq_data: dict, long_ex: CexManagerT, short_ex: CexManagerT) -> bool:
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
    sb_has_asset = f"{asset}{short_ex.EX_CURRENCY}".upper() in sb['open_position_set']
    if lb_has_asset and sb_has_asset:
        return True
    else:
        return False

def iexa_exit_pos(mq_data: dict, long_ex: CexManagerT, short_ex: CexManagerT) -> bool:
    if not iexa_check_pos(mq_data, long_ex, short_ex):
        # If there's no position to exit, discontinue
        PrettyColors().print_warning("Signaled asset NOT present, No orders sent")
        return False
    
    # Both market exit-position-order for an asset
    asset = mq_data["a"]
    lo_quantity = upbit_pos_balance(long_ex, f"{asset}".upper())
    long_ex.conn.create_order(
        f"{asset}/{long_ex.EX_CURRENCY}",
        "market",
        "sell",
        lo_quantity,
        mq_data["pl"]
    )
    # Add abs() to the quantity. Because short selling position is given to us in MINUS form.
    # For example, if you shorted 128 DOGE, it will return -128.
    so_quantity = abs(
        float(binance_pos_balance(short_ex, f"{asset}{short_ex.EX_CURRENCY}"))
    )
    short_ex.conn.create_order(
        f"{asset}/{short_ex.EX_CURRENCY}",
        "market",
        "buy",
        so_quantity,
        mq_data["ps"],
        params={"reduceOnly": True}
    )
    p = mq_data["pm"]
    report_order("exit", asset, lo_quantity, so_quantity, p)
    return True

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
