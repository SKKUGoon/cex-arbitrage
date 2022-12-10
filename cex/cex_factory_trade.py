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
    def order_buy(self, buy: dict):
        pass

    @abstractmethod
    def order_sell(self, sell: dict):
        pass

    @abstractmethod
    def order_tpsl(self, take_profit: float, stop_loss:float):
        # TPSL: Take Profit Stop Loss
        pass

    @abstractmethod
    def trade_result(self):
        pass


class CexFactoryT:
    def get_connection(self, exchange: CexManagerX):
        return exchange.connection()

    # Trader methods
    def get_balance(self, exchange: CexManagerT):
        return exchange.balance()

    def exec_order_sell(self, exchange: CexManagerT):
        return exchange.order_sell()

    def exec_order_buy(self, exchange: CexManagerT):
        return exchange.order_buy()

    def app_order_tpsl(self, take_profit: float, stop_loss:float, exchange:CexManagerT):
        return exchange.order_tpsl(
            take_profit=take_profit,
            stop_loss=stop_loss,
        )

    def report_trade_result(self, exchange: CexManagerT):
        return exchange