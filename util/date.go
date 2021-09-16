// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-12 2:46 下午
// version: 1.0.0
// desc   : 日期工具

package util

import (
	"strconv"
	"time"
)

// DatePattern ...
const (
	// 日期格式化模板
	DatePattern = "2006-01-02 15:04:05"
)

// 格式化日期
func Format(date time.Time) string {
	return date.Format(DatePattern)
}

// Now ...
func Now() time.Time {
	// 设置时区为：东八区
	return time.Now().In(time.FixedZone("CST", 8*3600))
}

// NowMillis ...
func NowMillis() int64 {
	return Now().UnixNano() / 1e6
}

// NowMillisStr ...
func NowMillisStr() string {
	return strconv.FormatInt(NowMillis(), 10)
}

// NowUnix ...
func NowUnix() int64 {
	return Now().Unix()
}

// NowUnixStr ...
func NowUnixStr() string {
	return strconv.FormatInt(NowUnix(), 10)
}
