# Initial Deploy Version

## Signal Maker

Signal maker is consisted of 2 parts. `1) Band updater` and `2) Near-Real-Time premium calculator`.

`1)` is in exec_band.py 
`2)` is in exec_pair_multi.py

Both will be inside a docker container. 

## Trader

TBD - after backtesting.

## Commits

`a1b2865` - Create non trade functions class factory `[exchangename]X` style for centralized exchange.

`7ae7fc2` - Create multiple dockerfiles for docker-compose.yml file. 
  - https://stackoverflow.com/questions/56432317/error-installing-gevent-in-docker-alpine-python for `gevent pip` error.
  - https://stackoverflow.com/questions/27409761/docker-multiple-dockerfiles-in-project for multiple dockerfiles in project information.

`b2ce6f8` - Create utility functions for trading and signal generating. Includes bollinger band, hedge ratio, leverage ratio calculation etc. 

`ba0af78` - Create trade signal generator. Mostly infinite loop that makes requests to webserver.

`777514c` - Create testing files for each class / functions. All testing should be included in `.dockerignore` as `test?`.

`615bd32` - Create operation necessary files.
  - in .gitignore file, all `.yaml` files are ignored - which should be remove because it contains private information.

`37a29bd` - Create `exec_` files. These files are the main files that each container executes

`1b6c9e8` - Create docker-compose files. Create/Update

`ebb1328` - Get Kimchi Premium backend from project kimchi or `cex-arb-back`. Attach it to main project

`6b95e2f` - Create healthcheck endpoint for kp-backend. Dockercompose file healthcheck purpose.

`728a23c` - Docker compose file, and Dockerfile (service building purpose) configuration. Binance Band and Binance WS are depending on trade-control container

`21ad1a5` - Creating trader class in python. - Websocket ability needed. 

`9dc20ab` - BUGFIX:: Python does not use `websocket` module. Instead it uses `websocket-client` module. Installing both with `pip install` cause major errors.

`43c13aa` - Add deployment IP address - (0.0.0.0) so that it's accessible.

`d234261`, `ec34061` `55108e5` - Add flag parse to employ deploy option. Add host name flag to specify docker container service name.

`7f2e579` - Add deploy environment (docker service name etc) to premium comparison program `exec_pair_multi.py`

`862a4f9` - Configure docker file. This includes container-wise networking and container healthcheck. Now executable. 

`03ef48b` - MODULE::TRADE:: Exchange balance search for trade prep.

`6f527f0` - Add key currency flag for both exchange. Python files. Change docker file accordingly

`29bd8c4` `bc52e9e` - Process absolute band threshold distance guard(1.5%point), for ensured profit during arbitrage. 

`d6532e3` - Change binance key currency from USDT to BUSD. Save transaction cost.

`84a1d64` - Change from redis cloud to redis local storage. Docker container. Check if do we need to open up external redis port?


`82a5645` - git add trader

`fff4c43` - Change trader alerting method from websocket to message queue. 

`f9e419a` - Delete unused functions. Single-websocket run etc.

`cd00191` - Delete old testfiles (TODO: Reminder: Write testfiles later again)

`f854845` - BUGFIX: host name should be service

`361ccde` - Refactor code: trade classes to *_trade.py files.

`cbac6b7` - Refactor code: Delete all the unused functions. + Pubsub method added. Publish signal message to redis pubsub. Move client to top of the function

`56f700b` - Example config file `exchange_example.yaml` added. 