from .coloring import PrettyColors
from typing import Dict


class Leverage:
    def __init__(self, wallet1_balance: float, wallet2_balance: float):
        self.wlt1 = wallet1_balance
        self.wlt2 = wallet2_balance
        self.wallets = {
            "wallet1": self.wlt1,
            "wallet2": self.wlt2,
        }

    def rcmd_leverage(self, fix_lev_rate_one: str, wallet1_usage: float, wallet2_usage: float) -> Dict:
        """
        Recommend proper leverage rate with wallet[standard] as the standard wallet.
        Standard wallet has a fixed leverage rate of one. Hence the name: `fix_lev_rate_one`
        For example, if 1) wallet1_usage is 100%, and wallet2_uage is 100%, 
        2) wallet1 is the standard; there can be 2 output result
        
        * wallet1 balance <= wallet2 balance: leverage is 1 : 1. 
        * wallet1 balance > wallet2 balance: leverage is 1 : n. 
          - To provide perfect hedging position, wallet2 should be leveraged.
          - For example, 
            wallet1: 100$ | wallet2: 50$. -> wallet2 should be x2 leveraged. 
        
        However, leverage should be integer. Provide the best integer for the leverage;
        which is the floor of the value. 
        """
        assert (0 <= wallet1_usage <= 1) and (0 <= wallet2_usage <= 1), (
            f"Usage must be less or equal than 1. Currently, w1: {wallet1_usage} w2: {wallet2_usage}"
        )
        assert fix_lev_rate_one in self.wallets.keys(), f"Unknown standard {fix_lev_rate_one}"

        if fix_lev_rate_one == "wallet1":
            lev = self._calc_leverage(
                self.wallets[fix_lev_rate_one] * wallet1_usage, 
                self.wallets["wallet2"] * wallet2_usage
            )
            return {"wallet1": lev[0], "wallet2": lev[1]}
        else:
            lev = self._calc_leverage(
                self.wallets[fix_lev_rate_one] * wallet2_usage, 
                self.wallets["wallet1"] * wallet1_usage
            )
            return {"wallet2": lev[0], "wallet1": lev[1]}
        
    
    @staticmethod
    def _calc_leverage(w_standard: float, w_compare: float) -> tuple:
        if w_compare == 0:
            return 0, 0

        if w_standard <= w_compare:
            return 1, 1
        else:
            return 1, w_standard // w_compare
                

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
