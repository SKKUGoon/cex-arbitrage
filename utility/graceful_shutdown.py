import sys


def handle_ctrlc(signum, stack_frame):
    sys.exit(0)
    