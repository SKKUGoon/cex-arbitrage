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

### Bugs and Fixes.
- In `requirements.txt` for python, <b>delete</b> `websocket` package. Python uses `websocket-client` package only. PIP Installing it both will cause errors. ( [9dc20ab](https://github.com/SKKUGoon/cex-arbitrage/commit/9dc20abf5a313665ea1f53325c7a4fe445218f57) )
- Delete unused functions such as single-websocket run, single asset band generation etc. Delete unused functions from backend, after changing from websocket method to redis message queue like model (PubSub) ( [f9e419a](https://github.com/SKKUGoon/cex-arbitrage/commit/f9e419a524fee5546b87cc16abece6ed0850bf17), [cd00191](https://github.com/SKKUGoon/cex-arbitrage/commit/cd00191c45c3ddd4f18b68dc316736557e946d79), [cbac6b7](https://github.com/SKKUGoon/cex-arbitrage/commit/cbac6b7bff50c7ad82827b613d9ba3ace312b817) )
- Refactor: script file name. Separate trading related classes into *_trade.py files ( [361ccde](https://github.com/SKKUGoon/cex-arbitrage/commit/361ccde9667bd0425dd677431bb31bfd41e03851) )
- Terminal not showing the work inside terminal even if its still working. Fix it by adding `flush=True` to all `print()` in python script. [Stackoverflow reference](https://stackoverflow.com/questions/74811707/python-docker-container-not-running-simultaneously/74811891#74811891). ( [a2dae2c] )