from utility.coloring import PrettyColors

from typing import Dict
import multiprocessing as mp
import json

import websocket


def binance_ws_multi(watch_list: dict, q_dict: dict, key_currency: str):
    # endpoint_url = "wss://stream.binance.com:9443/ws" <- for spot
    endpoint_url = "wss://fstream.binance.com/ws"
    my_id = 42

    def subscription(ticker_subject: dict) -> Dict:
        """
        @param ticker_subject: example {btcusdt: aggTrade}
        """
        PrettyColors().print_ok_cyan("binance websocket subscription message generated")
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
        PrettyColors().print_ok_cyan("binance opened")
        sub_msg = subscription(watch_list)
        ws.send(json.dumps(sub_msg))
    
    def on_message(ws, message, prev=None):
        data = json.loads(message)
        a = data['s'].split(key_currency.upper())[0]
        q_dict[a].put(data)
        # print("recv binance")

    def on_close(ws):
        PrettyColors().print_ok_cyan("binance closed")
        unsub_msg = unsubscription(watch_list)
        print(unsub_msg)
        ws.send(json.dumps(unsub_msg))
    
    ws = websocket.WebSocketApp(
        endpoint_url, 
        on_open=on_open,
        on_message=on_message,
        on_close=on_close,
    )
    ws.run_forever()
    ws.close()
