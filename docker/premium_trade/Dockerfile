FROM python:3.10.7

WORKDIR /app

# RUN apk add --no-cache python3-dev libffi-dev gcc musl-dev make

COPY requirements.txt .

RUN python -m pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT [ "python" ]
# why do we use busd only? to ensure low transaction fee
CMD [ "exec_trader.py", "-host=redis", "-upbitkey=KRW", "-binancekey=BUSD", "-upbitir=1", "-binanceir=1" ]