from utility.parse_yaml import ConfigParse


cp = ConfigParse("./binance.yaml")
d = cp.Parse()
print(d)
