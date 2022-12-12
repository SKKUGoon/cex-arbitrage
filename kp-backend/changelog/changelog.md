# cex-arb-backend. :: Changelog

## API Pre-Deployment.

`6e442de`
> Summary: On process of creating MySQL ORM.
> Basic Schemas are in place

`6d651ea`
> Summary: On process of creating redis ORM. 
> HSET and HGET ready. With test functions in place


`515f64d`
> Summary: At API start up, create gin Engine instance, 
> and redis Client.

`69c6bd4`
> Summary: Create API common tools. Such as CORS middleware.
> and Color coded print. 

`e67c3f1`
> Summary: Config file ignoring .gitconfig file. 

`bf61a18`
> Summary: Create handlers that gets band information. 
> TODO: create whether to trade or not decision function

`030a0bd`
> Summary: redis database commands

`52ca057`
> Summary: Websocket for server. Relays published information from redis.(Current)
> Change it to relaying trading information. 

`36d5226`
> Summary: Handle band input, premium POST request input. If premium input, the function compares
> it from the redis database. If condition is true, it relays trade order. 

`dafcbe1`
> Summary: Redis create op will return error if failed.

`7ba94be`
> Summary: Change pubsub subscribe from `channel1` to `channel_trade`. 
> TODO: pubsub channel name should be retrieved from Redis.yaml

`cd6a7c0`
> Summary: Process signal. Get premium - compare it with func, than relays premium 
> redis pubsub. 

`d175302`
> Summary: Redis message pubsub system. Publish Trading message, Subscribe to premium.

`edc5bea`
> Summary: Remove premium "\[get\]" method. remove wss from webserver.