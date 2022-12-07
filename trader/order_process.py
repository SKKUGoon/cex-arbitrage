import redis

r = redis.Redis(host="localhost", port=6379, db=0)
s = r.pubsub()

s.subscribe("channel_name")

while True:
    break


class Arbitrage: 
    def __init__(self):
        self.r = redis.Redis(host="localhost")
        self.topic: str = None

    def __conn__(self):
        ...

    def setup(self, value: str):
        self.topic = value
    
    async def handle_message(self):
        ...