from utility.coloring import PrettyColors
import time

def report_order(position: str, asset: str, long_qty: float, short_qty: float):
    msg = f"""{PrettyColors.BOLD}{PrettyColors.OKBLUE}Position {position}
    |  asset   | {asset}  
    | long Qty | {long_qty}
    | shrt Qty | {short_qty}
    |   time   | {time.strftime("%c")}{PrettyColors.ENDC}
    """
    print(msg)
