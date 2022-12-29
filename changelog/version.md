# Change Feature Logs for CEx-Arbitrage Repo

** commit id inside parenthesis is listed from the oldest to the most recent **

## Pre v0.5 (Non Released)
** This includes, v0.1, v0.2, v0.3, v0.4, and v0.5

### Main features
CEx-Arbitrage model with IEXA(<b>I</b>nter <b>EX</b>change <b>A</b>rbitrage) strategy can be divided into three main parts. `trade-control` webserver, `trader`, and `signal-maker`. First is made out of Go, and second and third is made out of python. In addition, <b>Redis</b> container is needed to store temporary band data. 

During the commits, it underwent a major design changes from using websockets to using message queue to deliver internal messages. ( [fff4c43](https://github.com/SKKUGoon/cex-arbitrage/commit/fff4c4300044af1586b6b1d12fc6e69601565045) ). Now new `changelog` logging style is being implemented. Starting from ( [123325b](https://github.com/SKKUGoon/cex-arbitrage/commit/123325b81705bf37c6455ae4aa2e69f967323835) ). 

- <b>Redis</b> Database. 
  - Changed from cloud to redis local storage inside docker container( [84a1d64](https://github.com/SKKUGoon/cex-arbitrage/commit/84a1d641343017b7c24edd0920f9e758bd473bd5) )

- `trade-control` features
  - Previously named Kimchi-Premium-Backend, or cex-arb-back, it is now attached to the main project. ( [ebb1328](https://github.com/SKKUGoon/cex-arbitrage/commit/ebb13281e464a8b9200c2d46d83e99e3fb3253d5) )
  - Add absoulte trading barrier of difference_premium > 1.5%point. This ensures the profitability of trade. ( [29bd8c4](https://github.com/SKKUGoon/cex-arbitrage/commit/29bd8c4f8d6fe1c89064f55e5c330a9254622690), [bc52e9e](https://github.com/SKKUGoon/cex-arbitrage/commit/bc52e9ec5903e3178c5eba01493a6571bee4c51f) )
  - Snippts of backend changelog can be found here. [cex-arb-back](https://github.com/SKKUGoon/cex-arb-back)

- `signal-maker` features
  - `[exchangename]X` styled class factory - for non trade functions. ( [a1b2865](https://github.com/SKKUGoon/cex-arbitrage/commit/a1b286549cda4b2021907c365addb39f4da6a46c) )
  - Trade signal generator. Previously in the form of `http` request, but now in the form of Publish to <b>Redis</b>. ( [ba0af78](https://github.com/SKKUGoon/cex-arbitrage/commit/ba0af78d8078bb8c9f50e50d088a7befbe7fe80a) )
  - Container restart process ( [a24582b](https://github.com/SKKUGoon/cex-arbitrage/commit/a24582b4a6e9f68f3eaee143bc3ef39356059713) )

- `trader` features
  - Trader classes are built with class factory styled. ( [21ad1a5](https://github.com/SKKUGoon/cex-arbitrage/commit/21ad1a5256f527b8efb5959e30ec9726ece3436e), [82a5645](https://github.com/SKKUGoon/cex-arbitrage/commit/82a56450cd179a248f710cc0e3dcd12236176add) )
  - Exchange balance API result value research. ( [03ef48b](https://github.com/SKKUGoon/cex-arbitrage/commit/03ef48bbdeff3ab65680fe6621a174b490cdae8c) )

- docker-compose
  - Individual dockerfiles for `docker-compose.yml` file. Create docker-compose files. After the backend attachment, create `healthcheck` for both backend and docker-compose file. Configure both `premium-band` container and `premium-signal` container depends on the health of `trade-control`. Configuration successful at commit id `862a4f9` ( [7ae7fc2](https://github.com/SKKUGoon/cex-arbitrage/commit/7ae7fc2480b70eefcbf9536094f3c11bd14fee50), 
  [1b6c9e8](https://github.com/SKKUGoon/cex-arbitrage/commit/1b6c9e8b21931a51fa1282934ba625a79a52691a), [6b95e2f](https://github.com/SKKUGoon/cex-arbitrage/commit/6b95e2f0719d60b810b576295509b74407ac00a4), 
  [728a23c](https://github.com/SKKUGoon/cex-arbitrage/commit/728a23ca5a5f4f67995925c45c21bcaa76c1f436), [862a4f9](https://github.com/SKKUGoon/cex-arbitrage/commit/862a4f9b5ecb2ac67a8fce11a010367624d5b453) )
    - https://stackoverflow.com/questions/56432317/error-installing-gevent-in-docker-alpine-python for `gevent pip` error.
    - https://stackoverflow.com/questions/27409761/docker-multiple-dockerfiles-in-project for multiple dockerfiles in project information.
  - Write python operation necessary files for docker. Script name starts with `exec_` ( [615bd32](https://github.com/SKKUGoon/cex-arbitrage/commit/615bd326a39aaf4ad897c6d8737fd52392d10dcd), [37a29bd](https://github.com/SKKUGoon/cex-arbitrage/commit/37a29bdd7208c80ecadf69881e132b18868bf9ed) )
  - Deploy address. IP address changed from 127.0.0.1 to 0.0.0.0. It's accessible from other containers.  ( [43c13aa](https://github.com/SKKUGoon/cex-arbitrage/commit/43c13aaab938ecabae52dec9cab5819c21492aee) )
  - Setup deploy environment and local environment discrimination for containers. ( [7f2e579](https://github.com/SKKUGoon/cex-arbitrage/commit/7f2e5791ecfa732b13c31ca6f021bab176643b9b) )

### Sub features
- Utility functions. Bollinger band, Calculating Hedge Ratio and Leerage Ratio. ( [b2ce6f8](https://github.com/SKKUGoon/cex-arbitrage/commit/b2ce6f86747ec56cceee2122e909aec6b04e35d5) )
- Test files for individual python scripts are created, then deleted. ( [777514c](https://github.com/SKKUGoon/cex-arbitrage/commit/777514c95a633ab2f91594a1a47c424efe7f35ad) )
- Add flag parse to employ deploy option. hostname, environment key asset etc. Add key currency flag for both exchange. ( [d234261](https://github.com/SKKUGoon/cex-arbitrage/commit/d234261c9e953b766a300abdc9240f2b6eda31f0), [ec34061](https://github.com/SKKUGoon/cex-arbitrage/commit/ec34061398e7d3eaf3e31f5db5e2657c81313175), [55108e5](https://github.com/SKKUGoon/cex-arbitrage/commit/55108e5521269bcf42100065e3e2f2a75c82d7d1), [6f527f0](https://github.com/SKKUGoon/cex-arbitrage/commit/6f527f0b78896416d160cc69c84f875f5b0ff8d0) )
- Trade `BUSD` not `USDT` since `BUSD` offers little-trading fees. ( [d6532e3](https://github.com/SKKUGoon/cex-arbitrage/commit/d6532e3597e3b7eeac4105d3a26f8535ea0e2ab0) )

### Bugs and Fixes
- In `requirements.txt` for python, <b>delete</b> `websocket` package. Python uses `websocket-client` package only. PIP Installing it both will cause errors. ( [9dc20ab](https://github.com/SKKUGoon/cex-arbitrage/commit/9dc20abf5a313665ea1f53325c7a4fe445218f57) )
- Delete unused functions such as single-websocket run, single asset band generation etc. Delete unused functions from backend, after changing from websocket method to redis message queue like model (PubSub) ( [f9e419a](https://github.com/SKKUGoon/cex-arbitrage/commit/f9e419a524fee5546b87cc16abece6ed0850bf17), [cd00191](https://github.com/SKKUGoon/cex-arbitrage/commit/cd00191c45c3ddd4f18b68dc316736557e946d79), [cbac6b7](https://github.com/SKKUGoon/cex-arbitrage/commit/cbac6b7bff50c7ad82827b613d9ba3ace312b817) )
- Refactor: script file name. Separate trading related classes into *_trade.py files ( [361ccde](https://github.com/SKKUGoon/cex-arbitrage/commit/361ccde9667bd0425dd677431bb31bfd41e03851) )
- Terminal not showing the work inside terminal even if its still working. Fix it by adding `flush=True` to all `print()` in python script. [Stackoverflow reference](https://stackoverflow.com/questions/74811707/python-docker-container-not-running-simultaneously/74811891#74811891). ( [a2dae2c]() )


## v0.6

### Main features
Add the features related to trading IEXA strategies. 


- BinanceFutureT, UpbitT (trader class) has `self.EX_CURRENCY` attribute. ( [a5e2a6d]() )
- Delete TPSL order (Take profit, stop loss) function in `CexManagerT` class and `CexFactoryT`. ( [84c0845]() )
- Create leverage calculation class `Leverage`( [ab464e4]() )
- Create flag for additional external variables - key currency for both exchange. Refactor Dockerfile. ( [ee999ba]() )
- Create callback function for entering, checking and exiting arbitrage position. ( [850525b]() )

### Sub features
- Delete unnecessary config files. Add Backend example configuration file. `config_example.yaml` (Not recorded on git) ( [ec7a949]() )
- CexManagerT class object's balance now returns open position asset set(). ( [6565d0b]() )

### Bugs and Fixes
- Fix wrong `time.sleep` calculation for container restart. Container restart after exit code (0) ( [23c2b03]() )


## v0.6.1

### Bugs and Fixes
- Fix aftermath of deleting config file `Redis.yaml` and `Config.yaml`.
  - Fix `Redis.yaml` disappearance. ./kp-backend/dao 's redis connection function, and yaml parsing tools. ( [e74f8b5]() )
  - Fix `Config.yaml` disappearance. ( [8c7d6b9]() )
  - Fix example `config_example.yaml` file ( [8c7d6b9]() )


## v0.7

### Main features
- Message Queue: Change trade_channel message. ( [3eee896]() )
  - Message now emits price information.
  - Changed from (giving out exchange-wise trade orders) to (trade pair-wise trade orders)
  - Price is delievered from python websocket `sig.py`. Transfers best ask price and best bid price from each exchange ( [cd5cf58]() )
- Trader Buying & Selling Now possible!  ( [b93ef6a]() )

### Sub features
- Hedge ratio 1:1 calculator function ( [ca02e4f]() )

### Bugs and Fixes
- Bugfix. Binance future trader class config file. API should have order authorization. ( [cee06ea]())


## v0.7.1

### Main features
Version v0.7.1 contains fixes after running field test1. Field test1 is conducted on December 21st, on POP-OS Linux machine. 

- Changes in docker compose file. ( [2855183]() )

### Bugs and Fixes
- Not trading Buy-sell pair, but was instead doing Buy-Buy. ( [e0a3c4d]() [b7f0b4d]() )
- Check balance again ( [5917d7b]() )


## v0.7.2

### Main features

Two bugs were discovered during 2 days of test run. 1) Not entering position with sell-buy but entering position in buy-buy. 2) Not able to detect already bought object. 1) is fixed in v0.7.1. Websocket problem is not python specific, but more efficient error-proof websocket module is being developed in branch `upgrade/wss`.

- 2) is fixed. Binance open_position_set needed keycurrency. For example, upbit has "WAVES" but in binance has "WAVESBUSD" ( [12de847]() )

### Bugs and Fixes
- `get_balance` after trading had missing variable input. Fixed. ( [412e74c]() )


## v0.7.3

### Main features
- Change websocket from spot stream to future stream ( [584f21d]() )
- Anticipate slippeage rate. Resize band width threshold from 0.015 to 0.02 ( [babaa34]() )

### Bugs and Fixes
- Fix `iexa_exit_pos`, by adding `abs()` to the quantity of the function. ( [a49e969]() )
  - Short position in binance will give you negative quantity(<0) value. 
  - When giving out market orders it should be ordering positive value.

## v0.8

### Main features
- Prep backend for new new strategy. ( [cb4762a]() )
  - Open a new message channel `notice_channel`
    - message example: `{"type": "notice", "status": true, "data": {"asset": "SOL", "complete": false}}`
  - Create structure for `notice_channel` message input. `Signal[BlockNotice]`
  - Add channel `NoticeMessage` in `SignalMessageQueue`
  - Organize goroutines
    - `mqToSigChan`: message queue signal_channel -> goroutine data channel
    - `mqToNotChan`: message queue notice_channel -> goroutine data channel
- New strategy - Notice Enclosed Premium ( [8e60c1a]() )
  - The premium on one exchange will rise, if there is not much coin influx. Take the premium by going long on the enclosed exchange, and hedging the position on the other exchange.
  - Crawler with python script. Dockerfile + Docker compose config file. Restart container every 1 minute. ( [27113a7]() )
  - If the message is old, (more then `RECENT_STANDARD` minute) it does not send any message to `notice_channel`. ( [412207b]() )

### Sub features
- Code refactoring
  - Create print_color functions ( [9733b45]() )
- `order_process.py`: add trade executed bool for `iexa_enter_pos` function and `iexa_exit_pos` function. If the return boolean is <b>False</b>, it does not update the balance - since no trade was made. Else if the return boolean is <b>True</b>, it updates the balance like before. This reduces calls to Binance and Upbit by not calling binance and upbit api everytime no-trade message is sent. ( [a79c691]() )  

### Bugs and Fixes
- Before: If signal in the message channel was malformed, it will emit msg to `trade_channel` regardless, which could lead to error in the trading module. But Now: if the signal message is malformed, it emits warning message in purple and continues loop. ( [cb4762a]() )
- Before: Position Exit signal could not be made if the band size shrink. But Now: Position Exit signal can be emitted regardless of the band size. ( [cb4762a]() )
- Before: no rules for printing colors in backend. But now: rules. ( [15ad222]() )
  - Backend ( [cb4762a]() )
    - Green Print : OK sign
    - Blue Print : Trade fail sign.
    - Cyan Print : Deploy environment status
    - Yellow Print : Any sort of operation. 
    - Purple Print : Warning or error. Not severe. Definitely not stopping to program.  
  - Python Scripts ( [ce5850f]() )
    - Bold Print : Reporting purpose
    - Underline Print : Infinite Loop Run start
    - Header Pring (Purple) : All CexManagerX, CexManagerT Function Order.
    - Green Print : OK sign
    - Blue Print : OK sign. But Not leading to trade. 
    - Cyan Print : OK sign - on Not Exchange Modules. (Signals)
    - Warning Print (Orange) : Unfavorable condition. But Continue. 
    - Fail Print (Red) : Outright Fail!ess)
- Delete password print in `CacheNewConn`. ( [cb4762a]() )
- Add `config_dev.yaml` in local storage (git ignored).
- Reduce `MINIMUM_BOUND_LENGTH` from 0.02 to 0.018. ( [cb4762a]() )

## v0.8.1

### Main features
- Add new field to Redis key-value database. ( [45b87af]() )
  - `beware` field: `asset`: `"no_enter"`
    - <b>ADDED</b>
      - `!p.Data.Complete` in `redis_pubsub.go` line 192
    - <b>DELETE</b> 
      - `p.Data.Complete` in `redis_pubsub.go` line 202

### Bugs and Fixes
- Delete unused directory `./kp-backend/ent` ( [e636339]() )
- Organize backend printed value. Now it looks like this. ( [e636339]() )
```console
2022/12/28 14:53:32  [| Asset: LINK    | Higher than thres 0.0058 < 0.0101 | BandSize: 0.004 |] 
2022/12/28 14:53:32  [| Asset: SAND    | Higher than thres 0.0057 < 0.0114 | BandSize: 0.004 |] 
```
- Add release mode according to environement. ( [e636339]() )
- Flagged Asset on notice. Bug in the handler. Logic update( [bc77150]() )


## v0.8.2

### Main features
- Tweak strategy. ( [b0d4156]() )
  - Adjust band generation data. from `5m` to `1d`. From `30` length to `20` length.
  - Band length to 4%point. 
  - Currently premium is at all time low (2022/12/30). Strategy change needed
- Change the hedging method. ( [d47d2ae]() )
  - From token quantity hedging to total spent capital hedging.
- Notice_channel -> trade_channel gives no input about price of the exchange. It gives you as value of `-1`. Normal ordering method will return error. ( [62cdc0a]() )
  - Upbit: Needs 1. quantity 2. price even if you are ordering through `market` price. They access `quantity * price` and buys for just that amount. 
  - Binance, if you order by `market`, `price` doesn't matter. 
  - For notice origin trading message, you must supply them with market price in order to hedge properly. 


## v0.8.3

### Main features
- Fix websocket app.
  - Make websocket app crash and restart on certain condition
    - When Binance App or upbit app automatically closes websocket connection, crash the container, and then restart. ( [dc0b86f]() )
      - Crashing is done by, setting `long_noval`, `short_noval` value. And at every iteration, if the websocket does not contain the asset's price value, add 1 respectively. If the `long_noval` or `short_noval` is equal to total number of asset, - which is highly unlikely since there are more than 20 assets, the app crashes.
      - Process Websocket is assigned in process 1 and process 2 multiprocess. They are set as daemon process, so that sys.exit() will forcefully close them.
      - Process 3 was originally inside multiprocess, but it's now taken out. 
      - Restarting is done by setting up docker compose file accordingly. ( [7028f00]() )
