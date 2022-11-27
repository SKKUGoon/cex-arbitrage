from .coloring import PrettyColors
from typing import Dict


class HedgeCalc:
    def __init__(self, leverage: int, trading_rules: dict) -> None:
        if not self._high_leverage(leverage, trading_rules):
            raise RuntimeError("leverage too high. Cancel order.")

        self.leverage = leverage
        
    def _high_leverage(lev: int, rules: dict) -> bool:
        if lev > rules["leverage"]:
            print(PrettyColors.FAIL + "Failed Calc. Leverage too high" + PrettyColors.ENDC)
            return False
        else:
            return True

    def calc(self, pair: dict) -> Dict:
        assert 'upbit' in pair.keys()
        assert 'binance' in pair.keys()

        if pair['upbit'] == "":
            # Binance is leveraged <self.leverage>
            # Upbit order quantity should be multiplied
            binance_q = float(pair['binance'])
            upbit_lev_q = binance_q * self.leverage

            new_pair = {
                "upbit": upbit_lev_q,
                "binance": pair["binance"]
            }
            return new_pair

        elif pair['binance'] == "":
            # Upbit order quantity should be 1/n multiplied
            upbit_q = float(pair['upbit'])
            binance_lev_q = upbit_q / self.leverage

            new_pair = {
                "upbit": pair["binance"],
                "binance": binance_lev_q
            }
            return new_pair
        
        else:
            raise RuntimeError("both pairs missing order quantity")
            