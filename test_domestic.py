from cex.cex_factory import CexFactoryX, CexFactoryT
from cex.domestic import UpbitX, UpbitT

cm = CexFactoryX()

upbit_x = UpbitX()

# Get_connection
cm.get_connection(upbit_x)

# Get_tradable
cm.get_tradable(upbit_x)

# # Get history
# d = cm.get_history(upbit_x)
# print(d)

tm = CexFactoryT()
upbit_t = UpbitT()

# Get Balance
bal = tm.get_balance(upbit_t)
print(bal)