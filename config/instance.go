// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-08-31 9:56
// version: 1.0.0
// desc   : 一些实例

package config

import (
	"github.com/yhyzgn/goat/config/yaml"
)

// Instance 配置实例
var (
	Instance = &Config{}
)

// LoadProfile 加载环境配置
func LoadProfile(filename string) error {
	return yaml.NewReader().Decode(filename, Instance)
}
