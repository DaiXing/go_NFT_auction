package util

// 检查错误。
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
