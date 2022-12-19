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
        print(
            PrettyColors.HEADER 
            + "Binance Future Trader Config file" 
            + PrettyColors.ENDC, 
            flush=True
        )
        cp = ConfigParse("./exchange.yaml")
        d = cp.parse()
        return {
            'apiKey': d[self.EX_ID]['info']['api-key'],
            'secret': d[self.EX_ID]['info']['api-pass'],
            'options': {'defaultType': 'future'}
        }
    
    def connection(self):
        # Create self.conn
        print(
            PrettyColors.HEADER 
            + "Binance Future Trader Connection" 
            + PrettyColors.ENDC,
            flush=True
        )
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
        print(
            PrettyColors.HEADER 
            + "Binance Future Account Balance" 
            + PrettyColors.ENDC,
            flush=True
        )
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
        print(
            PrettyColors.HEADER 
            + "Binance Future Order Buy Process" 
            + PrettyColors.ENDC,
            flush=True
        )
        # self.conn.create_order()
        return 

    def order_sell(self, sell: dict):
        print(
            PrettyColors.HEADER 
            + "Binance Future Order Sell Process" 
            + PrettyColors.ENDC,
            flush=True,
        )
        # self.conn.create_order()
        return
    
    def trade_result(self):
        print(
            PrettyColors.HEADER 
            + "Binance Future Trade Result" 
            + PrettyColors.ENDC,
            flush=True
        ) 
        return 