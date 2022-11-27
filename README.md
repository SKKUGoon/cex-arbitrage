# Kimchi Premium Trader

## Execute

### On local machine (Mac OS Ventura)

1. Install pyenv
2. Install python (3.10.7). `pyenv install 3.10.7`
3. Set as local python. `pyenv local 3.10.7`
4. Create virtual environment.  `python -m venv venv`

To execute virtual environment python:  `venv/bin/python`

To execute pip in virtual environment:  `venv/bin/pip`

### Major Jobs

1. exec_pair_multi.py: process websocket data - runs for 24 hours TOPs
2. exec_band.py: update band by five minutes. 

Create separate docker file for both of them using docker compose

### When creating python-alpine docker images

```
RUN apk add --no-cache python3-dev libffi-dev gcc musl-dev make
```

We need this line to install `gevent` inside `requirements.txt`.