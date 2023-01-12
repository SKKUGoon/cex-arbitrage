from utility.coloring import PrettyColors

from datetime import datetime, timedelta, timezone
from typing import Final
import requests
import json
import re

import redis

TAG: Final = "[안내]"
CLUE: Final = ["입출금", "일시", "중단"]
DONE: Final = "완료"
RECENT_STANDARD: Final = 1 # minute
NOTICE_CHANNEL = "notice_channel"

def gen_signal_notice(env: str, hostname: str, common_trade: set):
    PrettyColors().print_underline(val="GEN_SIGNAL_NOTICE INTERMITANT BURST")
    if env.lower() == "dev":
        r = redis.Redis(host=hostname, port=15214, db=0, password="mypassword")
    elif env.lower() == "deploy":
        r = redis.Redis(host=hostname, port=6379, db=0, password="mypassword")
    else:
        raise RuntimeError(f"environment env={env} is not one of the specified")
    
    # Get upbit notice board
    url = "https://api-manager.upbit.com/api/v1/notices"
    resp = requests.get(url)
    resp_data = resp.json()
    
    try:
        data = resp_data['data']["list"]
    except Exception as e:
        print(e)
        return
    
    for d in data:
        notice_ticker, is_complete, ok = _parse_title(d['title'])
        notice_channel = {"type": "notice", "status": True}
        
        if not is_complete and ok:
            dfmt = "%Y-%m-%dT%H:%M:%S%z"
            parsed_ct = datetime.strptime(d['created_at'], dfmt)
            if not _is_recent_info(parsed_ct, RECENT_STANDARD):
                PrettyColors().print_ok_blue(f"{', '.join(notice_ticker)} ENTER notice outdated.")
                continue
            # Enter position
            for t in notice_ticker:
                # If trade not supported: pass
                if t not in common_trade:
                    continue
                # Message Pub/Sub
                notice_channel["data"] = {"asset": t, "complete": False}
                packet_dump = json.dumps(notice_channel)
                r.publish(channel=NOTICE_CHANNEL, message=packet_dump)

        elif is_complete and ok:
            dfmt = "%Y-%m-%dT%H:%M:%S%z"
            parsed_ct = datetime.strptime(d['created_at'], dfmt)
            if not _is_recent_info(parsed_ct, RECENT_STANDARD):
                PrettyColors().print_ok_blue(f"{', '.join(notice_ticker)} EXIT notice outdated.")
                continue
            # Exit existing position
            notice_channel = {"type": "notice", "status": True}
            for t in notice_ticker:
                # If trade not supported: pass
                if t not in common_trade:
                    continue
                # Message Pub/Sub
                notice_channel["data"] = {"asset": t, "complete": True}
                packet_dump = json.dumps(notice_channel)
                r.publish(channel=NOTICE_CHANNEL, message=packet_dump)


def _is_recent_info(ct: datetime, how_recent: int) -> bool:
    now = datetime.now(timezone.utc)
    time_diff = now - ct
    if time_diff > timedelta(minutes=how_recent):
        return False
    return True

def _parse_title(t: str):
    """
    @param t: title string

    Find coin ticker inside the title string.
    Return [Ticker list, is_completed, is_valid_info]
    """
    # Sort out 
    if TAG not in t[:4].replace(" ", ""):
        return [""], False, False

    if any([c not in t for c in CLUE]):
        return [""], False, False

    try:
        # Single: 1. Get the coin ticker inside `( )`
        start_ = t.index("(")
        end_ = t.index(")")
        t_parse = t[start_ + 1 : end_]
        pattern = ("[^a-zA-Z0-9,]")
        t_parse = re.sub(pattern, "", t_parse)
        if t_parse == "":
            # Context inside `()` is not ticker. It could be DONE
            raise ValueError
        
    except ValueError as _:
        # Multi 2. Get the Big Alphabet.  
        pattern = ("[^a-zA-Z0-9,]")
        t_parse = re.sub(pattern, "", t)

    t_parse_multi = t_parse.split(",")
    if "" == t_parse_multi[0]:
        return [""], False, False

    if DONE in t:
        return t_parse_multi, True, True

    return t_parse_multi, False, True
