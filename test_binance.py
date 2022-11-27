from cex.cex_factory import CexFactoryX
from cex.binance import BinanceX

cm = CexFactoryX()
binance_x = BinanceX()

# Get_connection
cm.get_connection(binance_x)

# Get tradable
cm.get_tradable(binance_x)

# Get History
d = cm.get_history(binance_x)
print(d)

