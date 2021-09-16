// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-14 10:20 上午
// version: 1.0.0
// desc   : 一些语法小公举

package util

import (
	"strconv"
	"strings"
)

// 模拟三目运算符
func If(condition bool, positive, negative interface{}) interface{} {
	if condition {
		return positive
	}
	return negative
}

// ParseBool ...
func ParseBool(str string) (bool, error) {
	if str = strings.TrimSpace(str); str == "" {
		return false, nil
	}
	return strconv.ParseBool(str)
}

// Btoi ...
func Btoi(value bool) int {
	if value {
		return 1
	}
	return 0
}
