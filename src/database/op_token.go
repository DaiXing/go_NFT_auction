package database

func QueryToken(nftContract string, tokenId string) (*MockTokenPo, error) {
	var one MockTokenPo
	err1 := Db.Where("nft_contract = ?", nftContract).
		Where("token_id = ?", tokenId).Take(&one).Error
	return &one, err1
}

func ExistToken(nftContract string, tokenId string) bool {
	_, err := QueryToken(nftContract, tokenId)
	return err == nil
}

// 更新DB
func UpdateToken(nftContract string, tokenId string, updateFields map[string]any) error {
	one, err := QueryToken(nftContract, tokenId)
	if err == nil {
		// 用ID更新。
		err2 := Db.Model(&MockTokenPo{}).
			Where("id = ?", one.ID).
			UpdateColumns(updateFields).Error
		return err2
	}
	return err
}
