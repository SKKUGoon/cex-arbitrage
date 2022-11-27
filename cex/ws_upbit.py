from utility.coloring import PrettyColors

from typing import Dict
import multiprocessing as mp
import json
import uuid

import websocket


def upbit_ws(q: mp.Queue):
    endpoint_url = "wss://api.upbit.com/websocket/v1"
    my_id = 43

    def subscription(ticker_subject: dict) -> Dict:
        """
        @param ticker_subject: example {}
        """
        sub_key = list(ticker_subject.keys())[0]
        sub = [
            {"ticket": str(uuid.uuid4())[:6]},
            {
                "type": sub_key,
                "codes": ticker_subject[sub_key],
                'isOnlyRealtime': True,
            },
        ]
        return sub

    def on_open(ws):
        print(PrettyColors.UNDERLINE + "upbit opened" + PrettyColors.ENDC)
        subscribe_msg = subscription({
            "orderbook": ["KRW-BTC.1"],
        })
        ws.send(json.dumps(subscribe_msg))
    
    def on_message(ws, message, prev=None):
        data = json.loads(message.decode('utf-8'))
        q.put(data)

    def on_error(ws, msg):
        print(msg)

    def on_close(ws):
        print(PrettyColors.UNDERLINE + "upbit closed" + PrettyColors.ENDC)
    
    ws = websocket.WebSocketApp(
        endpoint_url,
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close,
    )
    ws.run_forever()


def upbit_ws_multi(watch_ls: dict, q_dict: dict):
    endpoint_url = "wss://api.upbit.com/websocket/v1"
    my_id = 43

    def subscription(ticker_subject: dict) -> Dict:
        """
        @param ticker_subject: example {}
        """
        print(PrettyColors.OKBLUE + "upbit websocket subscription message generated" + PrettyColors.ENDC)
        sub_key = list(ticker_subject.keys())[0]
        sub = [
            {"ticket": str(uuid.uuid4())[:6]},
            {
                "type": sub_key,
                "codes": ticker_subject[sub_key],
                'isOnlyRealtime': True,
            },
        ]
        return sub

    def on_open(ws):
        print(PrettyColors.UNDERLINE + "upbit opened" + PrettyColors.ENDC)
        sub_msg = subscription(watch_ls)
        ws.send(json.dumps(sub_msg))
    
    def on_message(ws, message, prev=None):
        data = json.loads(message.decode('utf-8'))
        a = data['code'].split("-")[1]
        q_dict[a].put(data)

    def on_error(ws, msg):
        print(msg)

    def on_close(ws):
        print(PrettyColors.UNDERLINE + "upbit closed" + PrettyColors.ENDC)
    
    ws = websocket.WebSocketApp(
        endpoint_url,
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close,
    )
    ws.run_forever()
