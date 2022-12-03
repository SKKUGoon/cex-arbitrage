from cex.cex_factory import CexFactoryX, CexFactoryT
from cex.binance_future import BinanceFutureX, BinanceFutureT

cm = CexFactoryX()
binance_x = BinanceFutureX()

# Get_connection
cm.get_connection(binance_x)

# Get tradable
cm.get_tradable(binance_x)

# # Get History
# d = cm.get_history(binance_x)
# # print(d)

tm = CexFactoryT()
binance_t = BinanceFutureT()

# Get Balance
bal = tm.get_balance(binance_t)
print(bal)

