from utility.fx import forex
from utility.bollinger import bollinger

s, _ = forex()
print(s)

test = [1, 2, 3, 4, 5]
l, h = bollinger(test)
print(l, h)