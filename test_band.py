from cex.binance_future import BinanceFutureX
from cex.domestic import UpbitX
from sig.sig0_gen import gen_band_iexa


a, b = gen_band_iexa({}, UpbitX(), BinanceFutureX())
print(a, b)
