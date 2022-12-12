from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory_trade import CexManagerT

from typing import Dict

import ccxt


class UpbitT(CexManagerT):
    def __init__(self) -> None:
        self.EX_ID = 'upbit'

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()

    def parse_yaml(self) -> Dict:
        # Create self.config
        print(PrettyColors.HEADER + "Upbit Config file" + PrettyColors.ENDC)
        cp = ConfigParse('./upbit.yaml')
        d = cp.parse()
        return {
            'apiKey': d[self.EX_ID]['info']['api-key'], 
            'secret': d[self.EX_ID]['info']['api-pass']
        }

    def connection(self):
        # Create self.conn
        print(PrettyColors.HEADER + "Upbit Connection" + PrettyColors.ENDC)
        conn = ccxt.upbit(config=self.config)
        return conn

    def balance(self, key_currency: str="KRW") -> Dict:
        print(PrettyColors.HEADER + "Upbit Account Balance" + PrettyColors.ENDC)
        b = self.conn.fetch_balance()
        # Upbit is a currency exchange.
        # Not explicitly a position        
        return {
            'key_balance': {
                'asset': key_currency,
                'balance': b[key_currency],
            },
            'open_position': [op for op in b['info'] if op['currency'] != key_currency]
        }

    def order_buy(self, buy: dict):
        print(PrettyColors.HEADER + "Upbit Order Buy Execute" + PrettyColors.ENDC)
        return 

    def order_sell(self, sell: dict):
        print(PrettyColors.HEADER + "Upbit Order Sell Execute" + PrettyColors.ENDC)
        return 

    def order_tpsl(self, take_profit: float, stop_loss: float):
        print(PrettyColors.HEADER + "Upbit Order TakeProfit StopLoss Execute" + PrettyColors.ENDC)
        return 

    def trade_result(self):
        print(PrettyColors.HEADER + "Upbit Future Trade Result" + PrettyColors.ENDC)
        return