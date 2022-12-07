FROM python:3.10.7

WORKDIR /app

# RUN apk add --no-cache python3-dev libffi-dev gcc musl-dev make

COPY requirements.txt .

RUN python -m pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT [ "python" ]
# why do we use busd only? to ensure low transaction fee
CMD [ "exec_pair_multi.py", "-env=deploy", "-host=trade_control", "-upbitkey=krw", "-binancekey=busd" ]