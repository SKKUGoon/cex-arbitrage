from cex.cex_factory import CexFactoryX
from cex.domestic import UpbitX

cm = CexFactoryX()

upbit_x = UpbitX()

# Get_connection
cm.get_connection(upbit_x)

# Get_tradable
cm.get_tradable(upbit_x)

# Get history
d = cm.get_history(upbit_x)
print(d)
