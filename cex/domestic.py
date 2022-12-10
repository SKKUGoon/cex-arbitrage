from utility.coloring import PrettyColors
from utility.parse_yaml import ConfigParse
from .cex_factory import CexManagerX
from .cex_factory_trade import CexManagerT

from typing import Dict

import ccxt


class UpbitX(CexManagerX):
    def __init__(self):
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
            'apiKey': d['info']['api-key'], 
            'secret': d['info']['api-pass']
        }

    def connection(self):
        # Create self.conn
        print(PrettyColors.HEADER + "Upbit Connection" + PrettyColors.ENDC)
        conn = ccxt.upbit(config=self.config)
        return conn

    @staticmethod
    def _key_currency(currency_ls: dict) -> Dict:
        """
        return {<key currency>: <supported currency>}
        """
        key_currency = dict()
        for c in currency_ls.keys():
            t = c.split("/")
            if t[1] not in key_currency.keys():
                key_currency[t[1]] = list()
            key_currency[t[1]].append(t[0])
        return key_currency

    def tradable(self):
        # Create self.curr
        print(PrettyColors.OKCYAN + "Upbit Tradables Update" + PrettyColors.ENDC)
        curr = self.conn.load_markets()
        key_curr_pair = self._key_currency(curr)
        return key_curr_pair
    
    def ticker(self, ticker_set: set, key_currency: str) -> Dict:
        result = {"orderbook": list()}
        for t in ticker_set:
            # .1 is there to ensure that we get only the best bid, best ask.
            result["orderbook"].append(
                f"{key_currency.upper()}-{t.upper()}.1"
            )
        return result

    def history(self, ticker: str, key_currency: str, hist_len: int=30):
        request_for = f"{ticker.upper()}/{key_currency.upper()}"
        hist = self.conn.fetch_ohlcv(request_for, '5m', limit=hist_len)
        hist = list(map(lambda row: row[4], hist))
        return hist


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
            'apiKey': d['info']['api-key'], 
            'secret': d['info']['api-pass']
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