from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory import CexManagerX, CexManagerT

from typing import Dict, List

import ccxt


class BinanceFutureX(CexManagerX):
    def __init__(self):
        self.EX_ID = 'binance_future'

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()

    def parse_yaml(self) -> Dict:
        # Create self.config
        print(PrettyColors.HEADER + "Binance Future Config file" + PrettyColors.ENDC)
        cp = ConfigParse("./binance.yaml")
        d = cp.parse()
        return {
            'apiKey': d['info']['api-key'], 
            'secret': d['info']['api-pass'],
            'options': {'defaultType': 'future'}
        }

    def connection(self):
        # Create self.conn
        print(PrettyColors.HEADER + "Binance Future Connection" + PrettyColors.ENDC)
        conn = ccxt.binance(config=self.config)
        return conn

    @staticmethod
    def _key_currency(currency_ls: dict) -> Dict:
        """
        return {<key currency>: <supported currency>}
        """
        key_currency = dict()
        for c in currency_ls.keys():
            t = c.split("/")
            if len(t) >= 2:
                if t[1] not in key_currency.keys():
                    key_currency[t[1]] = list()
            else:
                continue
            key_currency[t[1]].append(t[0])
        return key_currency

    def tradable(self) -> Dict:
        # Create self.curr
        print(PrettyColors.OKCYAN + "Binance Future Tradables Update" + PrettyColors.ENDC)
        curr = self.conn.load_markets()
        key_curr_pair = self._key_currency(curr)
        return key_curr_pair

    def ticker(self, ticker_set: set, key_curruency: str) -> Dict:
        result = dict()
        for t in ticker_set:
            result[f"{t}{key_curruency}".upper()] = ["bookTicker"] 
        return result

    def history(self, ticker: str, key_currency: str, hist_len: int=30) -> List:
        request_for = f"{ticker.upper()}/{key_currency.upper()}"
        hist = self.conn.fetch_ohlcv(request_for, '5m', limit=hist_len)
        # [1668702300000, 16583.03, 16621.95, 16574.67, 16616.88, 1128.21598]
        # [<Unix Time>, <open>, <high>, <low>, <close>, <volume>]
        # Process `hist` data. Get only close data.
        hist = list(map(lambda row: row[4], hist))
        
        return hist


class BinanceFutureT(CexManagerT):
    def __init__(self):
        self.EX_ID = 'binance_future'

        # Created by functions
        self.config = self.parse_yaml()
        self.conn = self.connection()
    
    def parse_yaml(self) -> Dict:
        # Create self.config
        print(PrettyColors.HEADER + "Binance Future Trader Config file" + PrettyColors.ENDC)
        cp = ConfigParse("./binance.yaml")
        d = cp.parse()
        return {
            'apiKey': d['info']['api-key'],
            'secret': d['info']['api-pass'],
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