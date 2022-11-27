from typing import Union
import numpy as np

def bollinger(values: list, band: int=2) -> Union[float, float]:
    lower = np.mean(values) + band * np.std(values)
    higher = np.mean(values) - band * np.std(values)
    return lower, higher