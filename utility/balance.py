class Wallet:
    def __init__(self, wallet_info: dict):
        self.wallet = wallet_info

    def can_order(self, price: float, is_floor: bool):
        """
        @ is_floor: if True, than delete the floor. 3.01 -> 3
        """
        pass

    def remain(self):
        pass