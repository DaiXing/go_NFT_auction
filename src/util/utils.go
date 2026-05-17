package util

import (
	"context"
	"fmt"
	"strings"
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

// 日志工具。
type LogMaker struct {
	lines []string
}

func (tt *LogMaker) AddKV(key string, value any) {
	valueStr := fmt.Sprint(value)
	line := key + " = " + valueStr
	tt.AddLine(line)
}
func (tt *LogMaker) AddLine(line string) {
	tt.lines = append(tt.lines, line)
}
func (tt *LogMaker) LogString() string {
	buf := strings.Builder{}
	buf.Grow(len(tt.lines) * 100)

	for _, line := range tt.lines {
		buf.WriteString(line)
		buf.WriteString("\n")
	}
	str := buf.String()
	Logger.Info(str)
	return str
}
