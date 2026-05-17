package database

// 查拍卖信息。
func QueryActionInfoByAuctionId(auctionId uint64) (*AuctionInfoPo, error) {
	var row AuctionInfoPo
	result := Db.Where("auction_id = ?", auctionId).Take(&row)
	return &row, result.Error
}

// 是否存在。
func ExistsAuctionId(auctionId uint64) bool {
	_, err := QueryActionInfoByAuctionId(auctionId)
	if err == nil {
		return true
	}
	return false
}

// 更新DB
func UpdateAuction(auctionId uint64, updateFields map[string]any) error {
	auction, err := QueryActionInfoByAuctionId(auctionId)
	if err == nil {
		// 用ID更新。
		err2 := Db.Model(&AuctionInfoPo{}).
			Where("id = ?", auction.ID).
			UpdateColumns(updateFields).Error
		return err2
	}
	return err
}
