// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-20 11:37
// version: 1.0.0
// desc   :

package main

import (
	"e-learn/config"
	"e-learn/engine"
	"e-learn/logger"
	"time"
)

func main() {
	// panic 降级处理
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
		time.Sleep(1 * time.Second)
	}()

	err := config.LoadProfile("resources/app.yml")
	if nil != err {
		panic(err)
	}

	// 初始化日志
	logger.Init(config.Instance.Logging)
	logger.Info("正启动于...")
	logger.Info("配置加载完成！")

	engine.New().Start()
}
