from .coloring import PrettyColors
import yaml

class ConfigParse:
    def __init__(self, file_addr: str) -> None:
        self.addr = file_addr

    def parse(self):
        with open(self.addr, "r") as file:
            try:
                data = yaml.safe_load(file)
                return data
            except yaml.YAMLError as exc:
                PrettyColors().print_fail(exc)

