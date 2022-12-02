# Kimchi Premium Trade

# Regarding Production

## Websocket

### A. `websocket` or `websocket-client`?

`websocket` module is wrong module. One should only use `websocket-client` module. 
Including `websocket` module would cause major errors.

# How to Execute

## Without Docker

### A. On local machine - `Python` codes (Mac OS Ventura)

1. Install pyenv
2. Install python (3.10.7). `pyenv install 3.10.7`
3. Set as local python. `pyenv local 3.10.7`
4. Create virtual environment.  `python -m venv venv`

To execute virtual environment python:  `venv/bin/python`

To execute pip in virtual environment:  `venv/bin/pip`

### B. On local machine - `Go` codes

`./kp-backend` codes are all written in golang. 

1. Install go (1.19)

```
go run .
```

## With Docker

>### When creating python-alpine docker images
>Due to the uniqueness of alpine OS, we need this line to install `gevent` inside `requirements.txt`.
>```
>RUN apk add --no-cache python3-dev libffi-dev gcc musl-dev make
>```

We will not be using python-alpine docker images for now. The size does not matter yet.
For production coding, implement flags for python execution. In order to change localhost, and deploy host (0.0.0.0)

### A. Use `docker-compose up`

# Major Files

1. exec_pair_multi.py: process websocket data - runs for 24 hours TOPs
2. exec_band.py: update band by five minutes. 
3. ./kp-backend: **BACKEND CODE WRITTEN IN GO**. Backend for trading. Pubsub and websocket
  - Originally from repo [cex-arb-back](https://github.com/SKKUGoon/cex-arb-back)

Create separate docker file for both of them using docker compose
To Depend one container to the other, `depends_on` key value should be added to `docker-compose.yaml` file.

```yaml
  premium_band:
    build: 
      context: .
      dockerfile: ./docker/premium_band/Dockerfile
    ports: 
      - "8081:8081"
    depends_on:
      trade_control:
        condition: service_healthy
```