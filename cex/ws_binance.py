from utility.coloring import PrettyColors

from typing import Dict
import multiprocessing as mp
import json

import websocket


def binance_ws(q: mp.Queue):
    endpoint_url = "wss://stream.binance.com:9443/ws"
    my_id = 42

    def subscription(ticker_subject: dict) -> Dict:
        """
        @param ticker_subject: example {btcusdt: aggTrade}
        """
        sub = {
            "method": "SUBSCRIBE",
            "params": list(),
            "id": my_id
        }

        for key, val in ticker_subject.items():
            for v in val:
                sub["params"].append(
                    f"{key.lower()}@{v}"
                )
        return sub

    def unsubscription(ticker_subject: dict) -> Dict:
        """
        Needed for graceful shutdown
        @param ticker_subject: example {btcusdt: aggTrade}
        """
        unsub = {
            "method": "UNSUBSCRIBE",
            "params": list(),
            "id": 42
        }
        for key, val in ticker_subject.items():
            for v in val:
                unsub["params"].append(
                    f"{key.lower()}@{v}"
                )
        return unsub

    # To use websocketApp - we need to install both 
    #   websocket and websocket-client.
    # Define websocketapp's callback function:
    #   on_open, on_message and on_close.
    def on_open(ws):
        print(PrettyColors.UNDERLINE + "binance opened" + PrettyColors.ENDC)
        subscribe_msg = subscription({
            "BTCUSDT": ["bookTicker"]
        })
        print(subscribe_msg)
        ws.send(json.dumps(subscribe_msg))
    
    def on_message(ws, message, prev=None):
        data = json.loads(message)
        q.put(data)
        # print("recv binance")

    def on_close(ws):
        print(PrettyColors.UNDERLINE + "binance closed" + PrettyColors.ENDC)
        unsubscribe_msg = unsubscription({
            "BTCUSDT": ["bookTicker"]
        })
        print(unsubscribe_msg)
        ws.send(json.dumps(unsubscribe_msg))
    
    ws = websocket.WebSocketApp(
        endpoint_url, 
        on_open=on_open,
        on_message=on_message,
        on_close=on_close,
    )
    ws.run_forever()


def binance_ws_multi(watch_list: dict, q_dict: dict):
    endpoint_url = "wss://stream.binance.com:9443/ws"
    my_id = 42

    def subscription(ticker_subject: dict) -> Dict:
        """
        @param ticker_subject: example {btcusdt: aggTrade}
        """
        print(PrettyColors.OKBLUE + "binance websocket subscription message generated", PrettyColors.ENDC)
        sub = {
            "method": "SUBSCRIBE",
            "params": list(),
            "id": my_id
        }

        for key, val in ticker_subject.items():
            for v in val:
                sub["params"].append(
                    f"{key.lower()}@{v}"
                )
        return sub

    def unsubscription(ticker_subject: dict) -> Dict:
        """
        Needed for graceful shutdown
        @param ticker_subject: example {btcusdt: aggTrade}
        """
        unsub = {
            "method": "UNSUBSCRIBE",
            "params": list(),
            "id": 42
        }
        for key, val in ticker_subject.items():
            for v in val:
                unsub["params"].append(
                    f"{key.lower()}@{v}"
                )
        return unsub

    # To use websocketApp - we need to install both 
    #   websocket and websocket-client.
    # Define websocketapp's callback function:
    #   on_open, on_message and on_close.
    def on_open(ws):
        print(PrettyColors.UNDERLINE + "binance opened" + PrettyColors.ENDC)
        sub_msg = subscription(watch_list)
        ws.send(json.dumps(sub_msg))
    
    def on_message(ws, message, prev=None):
        data = json.loads(message)
        a = data['s'].split("USDT")[0]
        q_dict[a].put(data)
        # print("recv binance")

    def on_close(ws):
        print(PrettyColors.UNDERLINE + "binance closed" + PrettyColors.ENDC)
        unsub_msg = unsubscription(watch_list)
        print(unsub_msg)
        unsubscribe_msg = unsubscription({
            "BTCUSDT": ["bookTicker"]
        })
        print(unsubscribe_msg)
        ws.send(json.dumps(unsubscribe_msg))
    
    ws = websocket.WebSocketApp(
        endpoint_url, 
        on_open=on_open,
        on_message=on_message,
        on_close=on_close,
    )
    ws.run_forever()
