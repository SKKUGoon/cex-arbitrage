import requests
from typing import Tuple


def forex() -> Tuple:
    """
    Function forex() returns foreign exchange data - namely
    KRW/1USD. Which is currently about 1350 at the moment of writing.
    It returns Tuple of (foreign exchange value, is_not_approximated)

    If the exchange value is not available for some reason, it will instead use
    approximated value - very conservative value of 1350, and it will return False 
    for a second value of the tuple - indicating that the value is approximated
    TODO: change hard coded 1350 to something that can be retrieved from redis.
    """
    URL = "https://quotation-api-cdn.dunamu.com/v1/forex/recent"
    query = {"codes": "FRX.KRWUSD"}
    fx = requests.get(URL, params=query)
    if fx.status_code == 200:
        data = fx.json()
        return data[0]["basePrice"], True
    else:
        return 1350.00, False
