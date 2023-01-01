from typing import Dict
from .fx import forex


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
                

def hedge_quantity(exchange_wo_lev_q: float, leverage: int) -> float:
    return exchange_wo_lev_q / leverage

def hedge_capital(exchange_wo_lev_m: float, other_exchange_prc: float):
    """
    No leverage needed. Because order goes in with AMOUNT of the token.
    This includes the leverage. One just use money 1/lev times.
    """
    fx, _ = forex()
    other_exchange_m = exchange_wo_lev_m / fx
    return other_exchange_m / other_exchange_prc
