# NFT拍卖的后台

拍卖平台，拍卖不同的NFT。

创建拍卖。设置NFT物品，时间范围，起拍价等。

参与拍卖。多人出价，价高者优先。

结束拍卖。如果有人出价，则成功，NFT物品给出价者，金钱给卖家；否则，失败。

## 结构

合约： https://github.com/DaiXing/solidity_NFT_auction

后台： https://github.com/DaiXing/go_NFT_auction

### 配置文件  

文件  project.yaml

包含数据库、以太坊、WEB、mock的配置。

### 数据库  

类型  mysql 

Token表  mock_token

拍卖表 auction_info

出价表 auction_bid

KV表 key_value

### 合约  

alchemy查NFT

NFT合约ABI

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

目录 mock ，包含NFT、拍卖的交易逻辑，触发事件入库，跑整个测试流程

## 运行

切换到项目目录

执行  go run main.go

## 单测

切换到目录 test

执行  go test -v

simple_test.go 测试查询功能

nft_auction_test.go 测试NFT和拍卖的全流程





