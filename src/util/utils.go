package util

import (
	"context"
	"time"
)

// 检查错误。
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewContext(timeoutSeconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
}
