from utility.coloring import PrettyColors
import time

def report_order(position: str, asset: str, long_qty: float, short_qty: float, premium_enter: float):
    msg = f"""{PrettyColors.BOLD}{PrettyColors.OKBLUE}Position {position}
    {PrettyColors.BOLD}| Asset    | {asset}  
    {PrettyColors.BOLD}| Long Qty | {long_qty}
    {PrettyColors.BOLD}| Shrt Qty | {short_qty}
    {PrettyColors.BOLD}| Premium  | {premium_enter}
    {PrettyColors.BOLD}| Time     | {time.strftime("%c")}{PrettyColors.ENDC}
    """
    print(msg)
