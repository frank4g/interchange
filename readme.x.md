- https://docs.starport.com/guide/interchange/02-init.html

## Build
cat.|instr.|output
-|-|-
```scaffold.chain(interchange)```|*```starport scaffold chain github.com/frank4g/interchange --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(dex)```|```starport scaffold module dex --ibc --ordering unordered --dep bank```|generated:```proto/dex/(tx,query,genesis,params)+packet.proto```<br>```x/dex/+module_ibc.go```<br>```testutil/keeper/dex.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```
```scaffold.map(sell-order-book)```|```starport scaffold map sell-order-book amountDenom priceDenom --no-message --module dex```
```scaffold.map(buy-order-book)```|```starport scaffold map buy-order-book amountDenom priceDenom --no-message --module dex```
```scaffold packet(create-pair)```|```starport scaffold packet create-pair sourceDenom targetDenom --module dex```
```scaffold packet(sell-order)```|```starport scaffold packet sell-order amountDenom amount:int priceDenom price:int --ack remainingAmount:int,gain:int --module dex```
```scaffold packet(buy-order)```|```starport scaffold packet buy-order amountDenom amount:int priceDenom price:int --ack remainingAmount:int,purchase:int --module dex```
```scaffold message(cancel-sell-order)```|```starport scaffold message cancel-sell-order port channel amountDenom priceDenom orderID:int --desc "Cancel a sell order" --module dex```
```scaffold message(cancel-buy-order)```|```starport scaffold message cancel-buy-order port channel amountDenom priceDenom orderID:int --desc "Cancel a buy order" --module dex```
```scaffold.map(denom-trace)```|```starport scaffold map denom-trace port channel origin --no-message --module dex```

## Changes(proto)
k|genesis|query|tx|packet|
-|-|-|-|-
```starport scaffold list```|
```starport scaffold map```|```repeated```|msg,rpc:```Response,Request```(Get,All)
```starport scaffold packet```|||rpc,Msg:```x```,```xReponse```|```PacketData```,```PacketAck```
```starport scaffold message```|||rpc,Msg:```x```,```xReponse```

## Added
k|keeper|event|types|codec
-|-|-|-|-
```starport scaffold packet```|```x.go```,```msg_server_x.go```|```eventTypeXPacket```|
```starport scaffold message```|```msg_server_x.go```||```message_x.go```|
## Flags
flag|action
-|-
```--no-message```|
```--indexed```|
```--ibc```|
```--ordering unordered```|
```--desc "..."```|
```--module x```|
```--dep bank```|

## Types
k|v
-|-
default|string
int|int32

## Test
k|v
-|-
```start mars app```|```starport chain serve -c mars.yml```
```start mars venus```|```starport chain serve -c venus.yml```
```creating an order book for a pair of tokens```|```interchanged tx dex send-create-pair dex channel-0 marscoin venuscoin --chain-id mars```
