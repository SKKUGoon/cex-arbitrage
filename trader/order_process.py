from cex.cex_factory_trade import CexManagerT, CexFactoryT
import redis


class ArbitrageIEXA:
    """
    Inter Exchange Arbitrage Strategy
    In this case,
      exchange_long: upbit
      exchange_short: binance
    """
    def __init__(self, exchange_long: CexManagerT, exchange_short: CexManagerT, hostname: str):
        self.exchange = CexFactoryT()
        self.long = exchange_long
        self.short = exchange_short

        r = redis.Redis(host=hostname, port=6379, db=0, password="mypassword")
        self.pubsub = r.pubsub()

    def listen(self, channel_name: str):
        self.pubsub.subscribe(channel_name)

        while True:            
            print('waiting for trading message')
            res = self.pubsub.get_message(timeout=5)
            if res is not None:
                print(f"message: {res}", flush=True)
        ...



    
