class PrettyColors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'
    
    def print_bold(self, val):
        print(self.BOLD + val + self.ENDC, flush=True)
    
    def print_underline(self, val):
        print(self.UNDERLINE + val + self.ENDC, flush=True)

    def print_status_purple(self, val):
        print(self.HEADER + val + self.ENDC, flush=True)
    
    def print_ok_blue(self, val):
        print(self.OKBLUE + val + self.ENDC, flush=True)
    
    def print_ok_green(self, val):
        print(self.OKGREEN + val + self.ENDC, flush=True)

    def print_ok_cyan(self, val):
        print(self.OKCYAN + val + self.ENDC, flush=True)

    def print_warning(self, val):
        print(self.WARNING + val + self.ENDC, flush=True)
    
    def print_fail(self, val):
        print(self.FAIL + val + self.ENDC, flush=True)

