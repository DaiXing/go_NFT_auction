package eth

import (
	"io"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"my.nft.auction/src/util"
)

// ABI
var abiJson string
var abiObj *abi.ABI

// 合约的地址。
var contractAddr common.Address

func InitContract() {
	addrHex := util.Params.Eth.AuctionContractAddress
	contractAddr = common.HexToAddress(addrHex)

	filename := "contract_abi.json"
	file, err := os.Open(filename)
	util.CheckError(err)
	defer file.Close()

	// 读json
	bytex, err2 := io.ReadAll(file)
	if err2 != nil {
		panic(err2)
	}
	abiJson = string(bytex)

	// 解析ABI
	abix, err3 := abi.JSON(strings.NewReader(abiJson))
	util.CheckError(err3)
	abiObj = &abix

	util.Logger.Info("Eth 初始化 contract 完成")
}
