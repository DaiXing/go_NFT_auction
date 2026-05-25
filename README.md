# NFT拍卖的后台

## 结构

合约： https://github.com/DaiXing/solidity_NFT_auction

后台： https://github.com/DaiXing/go_NFT_auction

### 配置文件  project.yaml

包含数据库、以太坊、WEB、mock的配置。

### 数据库  mysql 

拍卖表 auction_info

出价表 auction_bid

KV表 key_value

### 合约  

alchemy查NFT

拍卖合约ABI

扫描历史事件

监听实时事件

事件写mysql表

### WEB

查token列表API

查拍卖信息API

查出价信息API

查统计API

### Mock 

project.yaml 包含mock用的NFT合约地址、几个用户

目录 mock 包含NFT、拍卖的交易逻辑，触发事件入库

## 运行

切换到项目目录

执行  go run main.go

## 单测

切换到目录 test

执行  go test -v

simple_test.go 测试查询功能

nft_auction_test.go 测试NFT和拍卖的全流程





