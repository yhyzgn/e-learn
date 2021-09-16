package logger

import (
	"github.com/yhyzgn/gog"
)

// Logging 配置
type Logging struct {
	Level     string   `yaml:"level"`
	ShortFile bool     `yaml:"short-file"`
	Writer    []string `yaml:"writer"`
}

var (
	writerMap map[string]gog.Writer
)

func init() {
	writerMap = map[string]gog.Writer{
		"console": NewConsoleWriter(),
	}
	gog.Async(true)
	gog.CallSkip(2)
}

// Init ...
func Init(logging Logging) {
	gog.SetLevel(gog.ParseLevel(logging.Level))
	var writers []gog.Writer
	for _, name := range logging.Writer {
		writers = append(writers, writerMap[name])
	}
	gog.SetWriter(writers...)
	gog.ShortFile(logging.ShortFile)
}

// Trace 追踪打印
func Trace(value ...interface{}) {
	gog.Trace(value...)
}

// TraceTag 追踪打印
func TraceTag(tag string, value ...interface{}) {
	gog.TraceTag(tag, value...)
}

// TraceF 追踪打印
func TraceF(format string, args ...interface{}) {
	gog.TraceF(format, args...)
}

// TraceTagF 追踪打印
func TraceTagF(tag string, format string, args ...interface{}) {
	gog.TraceTagF(tag, format, args...)
}

// Debug 调试打印
func Debug(value ...interface{}) {
	gog.Debug(value...)
}

// DebugTag 调试打印
func DebugTag(tag string, value ...interface{}) {
	gog.DebugTag(tag, value...)
}

// DebugF 调试打印
func DebugF(format string, args ...interface{}) {
	gog.DebugF(format, args...)
}

// DebugTagF 调试打印
func DebugTagF(tag string, format string, args ...interface{}) {
	gog.DebugTagF(tag, format, args...)
}

// Info 普通信息打印
func Info(value ...interface{}) {
	gog.Info(value...)
}

// InfoTag 普通信息打印
func InfoTag(tag string, value ...interface{}) {
	gog.InfoTag(tag, value...)
}

// InfoF 普通信息打印
func InfoF(format string, args ...interface{}) {
	gog.InfoF(format, args...)
}

// InfoTagF 普通信息打印
func InfoTagF(tag string, format string, args ...interface{}) {
	gog.InfoTagF(tag, format, args...)
}

// Warn 警告打印
func Warn(value ...interface{}) {
	gog.Warn(value...)
}

// WarnTag 警告打印
func WarnTag(tag string, value ...interface{}) {
	gog.WarnTag(tag, value...)
}

// WarnF 警告打印
func WarnF(format string, args ...interface{}) {
	gog.WarnF(format, args...)
}

// WarnTagF 警告打印
func WarnTagF(tag string, format string, args ...interface{}) {
	gog.WarnTagF(tag, format, args...)
}

// Error 错误打印
func Error(value ...interface{}) {
	gog.Error(value...)
}

// ErrorTag 错误打印
func ErrorTag(tag string, value ...interface{}) {
	gog.ErrorTag(tag, value...)
}

// ErrorF 错误打印
func ErrorF(format string, args ...interface{}) {
	gog.ErrorF(format, args...)
}

// ErrorTagF 错误打印
func ErrorTagF(tag string, format string, args ...interface{}) {
	gog.ErrorTagF(tag, format, args...)
}

// Fatal 错误打印，并结束进程
func Fatal(value ...interface{}) {
	gog.Fatal(value...)
}

// FatalTag 错误打印，并结束进程
func FatalTag(tag string, value ...interface{}) {
	gog.FatalTag(tag, value...)
}

// FatalF 错误打印，并结束进程
func FatalF(format string, args ...interface{}) {
	gog.FatalF(format, args...)
}

// FatalTagF 错误打印，并结束进程
func FatalTagF(tag string, format string, args ...interface{}) {
	gog.FatalTagF(tag, format, args...)
}
