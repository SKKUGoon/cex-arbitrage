from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory_trade import CexManagerT

from typing import Dict, List

import ccxt


class BinanceFutureT(CexManagerT):
    def __init__(self):
        self.EX_ID = 'binance'

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()
    
    def parse_yaml(self) -> Dict:
        # Create self.config
        print(PrettyColors.HEADER + "Binance Future Trader Config file" + PrettyColors.ENDC)
        cp = ConfigParse("./binance.yaml")
        d = cp.parse()
        return {
            'apiKey': d[self.EX_ID]['info']['api-key'],
            'secret': d[self.EX_ID]['info']['api-pass'],
            'options': {'defaultType': 'future'}
        }
    
    def connection(self):
        # Create self.conn
        print(PrettyColors.HEADER + "Binance Future Connection" + PrettyColors.ENDC)
        conn = ccxt.binance(config=self.config)
        return conn

    def open_position(self, total_position: list) -> List:
        opened = list()
        for p in total_position:
            if float(p['positionAmt']) != 0:
                opened.append(p)
        return opened
        
    def balance(self, key_currency: str="USDT") -> Dict:
        print(PrettyColors.HEADER + "Binance Future Account Balance" + PrettyColors.ENDC)
        # balance: { 
        #   'asset': 'USDT', 
        #   'balance': { 
        #       'free': ..., 
        #       'used': ..., 
        #       'total': ..., 
        #   } 
        # },
        # open_position: [
        #   { open position infos ... },
        # ]
        b = self.conn.fetch_balance()
        return {
            'key_balance': {
                'asset': key_currency,
                'balance': b[key_currency],
            },
            'open_position': self.open_position(b['info']['positions']),
        }

    def order_buy(self, buy: dict):
        print(PrettyColors.HEADER + "Binance Future Order Buy Execute" + PrettyColors.ENDC)
        # self.conn.create_order()
        return 

    def order_sell(self, sell: dict):
        print(PrettyColors.HEADER + "Binance Future Order Sell Execute" + PrettyColors.ENDC)
        # self.conn.create_order()
        return 
    
    def order_tpsl(self, take_profit: float, stop_loss: float):
        print(PrettyColors.HEADER + "Binance Future TakeProfit StopLoss Order" + PrettyColors.ENDC)
        return 
    
    def trade_result(self):
        print(PrettyColors.HEADER + "Binance Future Trade Result" + PrettyColors.ENDC) 
        return 