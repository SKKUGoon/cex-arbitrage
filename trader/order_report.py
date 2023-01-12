from utility.coloring import PrettyColors
import time

def report_order(position: str, asset: str, long_prc: float, short_prc: float, long_qty: float, short_qty: float, fx: float, leverage: int, premium_enter: float):
    msg = f"""{PrettyColors.BOLD}{PrettyColors.OKBLUE}Position {position}
    {PrettyColors.BOLD}+==========+==================================
    {PrettyColors.BOLD}| Asset    | {asset} with leverage {leverage}
    {PrettyColors.BOLD}+----------+----------------------------------
    {PrettyColors.BOLD}| Quantity | 
    {PrettyColors.BOLD}| Long Qty | {long_qty}
    {PrettyColors.BOLD}| Shrt Qty | {short_qty}
    {PrettyColors.BOLD}+----------+----------------------------------
    {PrettyColors.BOLD}| Spent    |
    {PrettyColors.BOLD}| Long Exc | {long_prc} KRW
    {PrettyColors.BOLD}| Shrt Exc | {short_prc * short_qty} USD ({short_prc * fx} KRW, fx adjusted)
    {PrettyColors.BOLD}| SpntCalc | {short_prc * short_qty * leverage * fx} KRW Shrt ~ {long_prc} KRW long
    {PrettyColors.BOLD}+----------+----------------------------------
    {PrettyColors.BOLD}| Premium  | {premium_enter}
    {PrettyColors.BOLD}| Time     | {time.strftime("%c")}{PrettyColors.ENDC}
    {PrettyColors.BOLD}+==========+==================================
    """
    print(msg)
