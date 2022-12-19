from cex.cex_factory_trade import CexFactoryT, CexManagerT
from utility.hedge import HedgeCalc
from utility.coloring import PrettyColors


def iexa_enter_pos(mq_data: dict, lev: int, balance: dict, order_ratio: float, long_ex: CexManagerT, short_ex: CexManagerT):
    if iexa_check_pos(mq_data, long_ex, short_ex):
        # If asset position is already present, discontinue
        return

    # Both market enter-position-order for an asset
    if "data" not in mq_data.keys():
        return
    asset = mq_data["data"]["a"]

    lo_money = balance['l'] * order_ratio
    so_money = balance['s'] * order_ratio

    long_ex.conn.create_market_buy_order(
        symbol=f"{asset}/{long_ex.EX_CURRENCY}",
        amount="",
    )
    # Adjust leverage for IEXA arb
    short_ex.fapiPrivate_post_leverage({
        "symbol": f"{asset}/{short_ex.EX_CURRENCY}",
        "leverage": lev
    })
    short_ex.conn.create_market_sell_order(
        symbol=f"{asset}/{short_ex.EX_CURRENCY}",
        amount="",
    )

def iexa_check_pos(mq_data: dict, long_ex: CexManagerT, short_ex: CexManagerT):
    """
    Each `CexManagerT` object has method `balance`. `balance` returns 
    "open_position_set" keys.

    return True: If asset in both long_exchange and short_exchange balance
    return False: Else
    """
    if "data" not in mq_data.keys():
        return
    asset = mq_data["data"]["a"]

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
    if "data" not in mq_data.keys():
        return
    asset = mq_data["data"]["a"]
    
    long_ex.conn.create_market_sell_order(
        symbol=f"{asset}/{long_ex.EX_CURRENCY}",
        amount="",
    )
    short_ex.conn.create_market_buy_order(
        symbol=f"{asset}/{short_ex.EX_CURRENCY}",
        amount="",
        params={"reduceOnly": True}
    )
 