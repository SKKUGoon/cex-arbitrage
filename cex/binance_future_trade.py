from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory_trade import CexManagerT

from typing import Dict

import ccxt


class BinanceFutureT(CexManagerT):
    def __init__(self, key_currency: str):
        self.EX_ID = 'binance'
        self.EX_CURRENCY = key_currency.upper()

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()

    def parse_yaml(self) -> Dict:
        # Create self.config
        PrettyColors().print_status_purple("ESSENTIAL: Binance Future Trader Config file")
        cp = ConfigParse("./exchange.yaml")
        d = cp.parse()
        return {
            'apiKey': d[self.EX_ID]['trade']['api-key'],
            'secret': d[self.EX_ID]['trade']['api-pass'],
            'options': {'defaultType': 'future'}
        }
    
    def connection(self):
        # Create self.conn
        PrettyColors().print_status_purple("ESSENTIAL: Binance Future Trader Connection")
        conn = ccxt.binance(config=self.config)
        return conn

    @staticmethod
    def _open_position(total_position: list) -> tuple:
        opened = list()
        opened_set = set()
        for p in total_position:
            if float(p['positionAmt']) != 0:
                opened.append(p)
                opened_set.add(p['symbol'])
        return opened, opened_set
        
    def balance(self, key_currency: str) -> Dict:
        PrettyColors().print_status_purple("Binance Future Account Balance")
        # balance: {... 'balance': {'free': ..., 'used': ..., 'total': ...,}} 
        # open_position: [{ open position infos ... }]
        b = self.conn.fetch_balance()
        o_pos, o_pos_set = self._open_position(b['info']['positions'])
        return {
            'key_balance': {
                'asset': key_currency,
                'balance': b[key_currency],
            },
            'open_position': o_pos,
            'open_position_set': o_pos_set,
        }

    def order_buy(self, buy: dict):
        PrettyColors().print_status_purple("Binance Future Order Buy Process")
        # self.conn.create_order()
        return 

    def order_sell(self, sell: dict):
        PrettyColors().print_status_purple("Binance Future Order Sell Process")
        # self.conn.create_order()
        return
    
    def trade_result(self):
        PrettyColors().print_status_purple("Binance Future Trade Result")
        return 