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

`ebb1328` - Get Kimchi Premium backend from project kimchi or `cex-arb-back`