// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-06 15:19
// version: 1.0.0
// desc   :

package logger

import (
	"io"
	"os"

	"github.com/yhyzgn/gog"
)

// ConsoleWriter 控制台输出器
type ConsoleWriter struct {
	out io.WriteCloser
}

// NewConsoleWriter 创建控制台输出器对象
func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		out: os.Stdout,
	}
}

// Write 输出日志
func (cw *ConsoleWriter) Write(info *gog.LogInfo, data []byte) (n int, err error) {
	// 忽略请求日志
	return cw.out.Write(data)
}

// Close 关闭输出流
func (cw *ConsoleWriter) Close() error {
	return cw.out.Close()
}
