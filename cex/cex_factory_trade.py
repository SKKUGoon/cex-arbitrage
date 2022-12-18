from abc import ABC, abstractmethod
from .cex_factory import CexManagerX


class CexManagerT(ABC):
    @abstractmethod
    def connection(self):
        pass

    @abstractmethod
    def balance(self, key_currency: str):
        pass

    @abstractmethod
    def order_buy(self, buy: dict, leverage: int):
        pass

    @abstractmethod
    def order_sell(self, sell: dict):
        pass

    @abstractmethod
    def trade_result(self):
        pass


class CexFactoryT:
    def get_connection(self, exchange: CexManagerX):
        return exchange.connection()

    # Trader methods
    def get_balance(self, exchange: CexManagerT, key: str):
        return exchange.balance(key.upper())

    def exec_order_buy(self, exchange: CexManagerT):
        return exchange.order_buy()

    def exec_order_sell(self, exchange: CexManagerT):
        return exchange.order_sell()
        
    def report_trade_result(self, exchange: CexManagerT):
        return exchange