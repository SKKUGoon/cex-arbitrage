from utility.coloring import PrettyColors
import time

def report_order(position: str, asset: str, long_qty: float, short_qty: float) -> str:
    msg = f"""Position {position}
    |  asset   | {asset}  
    | long Qty | {long_qty}
    | shrt Qty | {short_qty}
    |   time   | {time.strftime("%c")}
    """

    print(
        PrettyColors.BOLD + PrettyColors.OKBLUE
        + msg
        + PrettyColors.ENDC
    )

    return ""