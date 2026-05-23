package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"my.nft.auction/src/util"
)

// 用户信息。包含私钥、公钥、地址。
type UserInfo struct {
	PrivateKeyHex string
	PrivateKey    *ecdsa.PrivateKey `json:"-"`
	PublicKey     *ecdsa.PublicKey
	Addr          common.Address `json:"-"`
	AddrHex       string
}

// 用私钥，生成1个user
func UserFromPrivateKey(privateKeyHex string) *UserInfo {
	// 去掉前缀。
	privateKeyHex2 := privateKeyHex[2:]
	// 解析私钥。
	privateKey, err1 := crypto.HexToECDSA(privateKeyHex2)
	util.CheckError(err1)
	// 公钥
	pubKey := privateKey.PublicKey
	// 地址。
	addr := crypto.PubkeyToAddress(pubKey)

	user := &UserInfo{
		PrivateKeyHex: privateKeyHex,
		PrivateKey:    privateKey,
		PublicKey:     &pubKey,
		Addr:          addr,
		AddrHex:       addr.Hex(),
	}
	return user
}
