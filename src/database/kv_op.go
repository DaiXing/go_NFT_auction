package database

// 查询。
func QueryKeyValuePo(key string) (*KeyValuePo, error) {
	var kv KeyValuePo
	err := Db.Where("param_key = ?", key).Take(&kv).Error
	return &kv, err
}

// 查值。 带默认值。
func QueryKeyValue(key string, valueDefault string) string {
	kv, err := QueryKeyValuePo(key)
	if err != nil {
		return valueDefault
	}
	return kv.ParamValue
}

// 更新。
func UpdateKeyValue(key string, value string) error {
	kv, err := QueryKeyValuePo(key)
	// 更新。
	if err == nil {
		err2 := Db.Model(&KeyValuePo{}).
			Where("id = ?", kv.ID).
			UpdateColumns(map[string]interface{}{"param_value": value}).Error
		return err2
	} else {
		// 新增。
		kv2 := KeyValuePo{
			ParamKey:   key,
			ParamValue: value,
		}
		return Db.Create(&kv2).Error
	}
}
