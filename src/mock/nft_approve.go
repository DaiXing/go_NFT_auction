package mock

import "my.nft.auction/src/eth"

func InitNftApprove() {
	// 直接授权全部。简化测试。
	caller := UserJack
	toAddr := eth.AuctionContractAddr
	TxNftApproveAll(caller, toAddr)
}
