- https://docs.starport.com/guide/ibc.html

## Build
cat.|instr.|output
-|-|-
```scaffold.chain(interchange)```|*```starport scaffold chain github.com/frank4g/interchange --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(dex)```|```starport scaffold module dex --ibc --ordering unordered --dep bank```|generated:```proto/dex/(tx,query,genesis,params)+packet.proto```<br>```x/dex/+module_ibc.go```<br>```testutil/keeper/dex.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```
```starport scaffold map sell-order-book amountDenom priceDenom --no-message --module dex```
## Changes
k|genesis.proto|query.proto
-|-
```starport scaffold list```|
```starport scaffold map```|```repeated```|msg,rpc:```Response,Request```(Get,All)
