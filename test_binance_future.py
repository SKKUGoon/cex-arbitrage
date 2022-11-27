from cex.cex_factory import CexFactoryX
from cex.binance_future import BinanceFutureX

cm = CexFactoryX()
binance_x = BinanceFutureX()

# Get_connection
cm.get_connection(binance_x)

# Get tradable
cm.get_tradable(binance_x)

# Get History
d = cm.get_history(binance_x)
print(d)