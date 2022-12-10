from abc import ABC, abstractmethod


class CexManagerX(ABC):
    @abstractmethod
    def parse_yaml(self):
        pass

    @abstractmethod
    def connection(self):
        pass
    
    @abstractmethod
    def tradable(self):
        pass

    @abstractmethod
    def ticker(self, ticker_set: set, key_currency: str):
        pass

    @abstractmethod
    def history(self, ticker: str, key_currency: str, hist_len: int=30):
        pass


class CexFactoryX:
    # Exchange methods
    def get_connection(self, exchange: CexManagerX):
        return exchange.connection()

    def get_tradable(self, exchange: CexManagerX):
        return exchange.tradable()

    def send_signal(self, exchange: CexManagerX):
        send_to = ""
        return exchange.signal(send_to)

    def gen_ticker(self, exchange: CexManagerX, ticker_set: set, key_curruency: str):
        return exchange.ticker(ticker_set, key_curruency)

    def get_history(self, exchange: CexManagerX, ticker: str, key_currency: str):
        return exchange.history(ticker, key_currency) 
        

