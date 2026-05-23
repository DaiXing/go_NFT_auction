package database

// 查 竞拍。
func QueryBidByBidId(bidId string) (*AuctionBidPo, error) {
	var row AuctionBidPo
	err := Db.Where("bid_id = ?", bidId).Take(&row).Error
	return &row, err
}

// 是否存在
func ExistsBidId(bidId string) bool {
	_, err := QueryBidByBidId(bidId)
	if err == nil {
		return true
	}
	return false
}

// 更新。
func UpdateBid(bidId string, updateFields map[string]any) error {
	bidInfo, err := QueryBidByBidId(bidId)
	if err == nil {
		// 用ID更新。
		err2 := Db.Model(&AuctionBidPo{}).
			Where("id = ?", bidInfo.ID).
			UpdateColumns(updateFields).Error
		return err2
	}
	return err
}
