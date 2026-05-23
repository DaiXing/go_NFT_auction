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

	// 解析ABI
	abix, err3 := AbiFromFile("ABI_auction.json")
	util.CheckError(err3)
	abiObj = abix

	util.Logger.Info("Eth 初始化 contract ABI 完成")
}

// 用json文件生成ABI
func AbiFromFile(filename string) (*abi.ABI, error) {
	logMaker := util.LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine("用json文件生成ABI")
	logMaker.AddKV(" filename", filename)

	file, err := os.Open(filename)
	if err != nil {
		logMaker.AddKV(" os.Open", err)
		return nil, err
	}
	defer file.Close()

	// 读json
	bytex, err2 := io.ReadAll(file)
	if err2 != nil {
		logMaker.AddKV(" io.ReadAll", err2)
		return nil, err2
	}
	logMaker.AddKV(" 文件字节数", len(bytex))
	abiJson = string(bytex)

	// 解析ABI
	abix, err3 := abi.JSON(strings.NewReader(abiJson))
	if err3 != nil {
		logMaker.AddKV(" abi.JSON", err3)
		return nil, err3
	}
	return &abix, nil
}
