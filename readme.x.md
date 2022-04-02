- https://docs.starport.com/guide/ibc.html

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

## Changes(proto)
k|genesis|query|tx|packet|
-|-
```starport scaffold list```|
```starport scaffold map```|```repeated```|msg,rpc:```Response,Request```(Get,All)
```starport scaffold packet```|-|-|rpc,Msg:```x```,```xReponse```|```PacketData```,```PacketAck```

## Added
k|keeper|event|
-|-
```starport scaffold packet```|x.go,msg_server_x.go|```eventTypeXPacket```|

## Flags
flag|action
-|-
```--no-message```|
```--indexed```|
```--ibc```|
```--ordering unordered```|
```--module x```|
```--dep bank```|

## Types
k|v
-|-
default|string
int|int32
