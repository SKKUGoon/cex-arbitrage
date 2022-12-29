from utility.coloring import PrettyColors
import time

def report_order(position: str, asset: str, long_qty: float, short_qty: float, fx: float, premium_enter: float):
    msg = f"""{PrettyColors.BOLD}{PrettyColors.OKBLUE}Position {position}
    {PrettyColors.BOLD}| Asset    | {asset}  
    {PrettyColors.BOLD}| Long Prc | {long_qty} won
    {PrettyColors.BOLD}| Shrt Prc | {short_qty} usd * {fx} krw/usd = {short_qty * fx} won
    {PrettyColors.BOLD}| Premium  | {premium_enter}
    {PrettyColors.BOLD}| Time     | {time.strftime("%c")}{PrettyColors.ENDC}
    """
    print(msg)
