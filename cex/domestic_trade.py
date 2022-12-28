from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory_trade import CexManagerT

from typing import Dict

import ccxt


class UpbitT(CexManagerT):
    def __init__(self, key_currency: str):
        self.EX_ID = 'upbit'
        self.EX_CURRENCY = key_currency.upper()

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()

    def parse_yaml(self) -> Dict:
        # Create self.config
        PrettyColors().print_status_purple("ESSENTIAL: Upbit Trader Config file")
        cp = ConfigParse("./exchange.yaml")
        d = cp.parse()
        return {
            'apiKey': d[self.EX_ID]['info']['api-key'], 
            'secret': d[self.EX_ID]['info']['api-pass']
        }

    def connection(self):
        # Create self.conn
        PrettyColors().print_status_purple("ESSENTIAL: Upbit Trader Connection")
        conn = ccxt.upbit(config=self.config)
        return conn

    @staticmethod
    def _open_position(total_position: list, key_currency: str) -> tuple:
        opened = list()
        opened_set = set()
        for p in total_position:
            if p['currency'] != key_currency:
                opened.append(p)
                opened_set.add(p['currency'])
        return opened, opened_set

    def balance(self, key_currency: str="KRW") -> Dict:
        PrettyColors().print_status_purple("Upbit Account Balance")
        b = self.conn.fetch_balance()
        # Upbit is a currency exchange.
        # Not explicitly a position    
        o_pos, o_pos_set = self._open_position(b['info'], key_currency)    
        return {
            'key_balance': {
                'asset': key_currency,
                'balance': b[key_currency],
            },
            # 'open_position': [op for op in b['info'] if op['currency'] != key_currency],
            'open_position': o_pos,
            'open_position_set': o_pos_set,
        }

    def order_buy(self, buy: dict):
        PrettyColors().print_status_purple("Upbit Order Buy Process")
        return 

    def order_sell(self, sell: dict):
        PrettyColors().print_status_purple("Upbit Order Sell Process")
        return 

    def trade_result(self):
        PrettyColors().print_status_purple("Upbit Trade Result")
        return