// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-06-06 02:58
// version: 1.0.0
// desc   :

package orm

import (
	"e-learn/config"
	"e-learn/logger"
	"e-learn/util"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var (
	once sync.Once
	DB   *gorm.DB
)

// Connect 连接到 mysql 服务
func Connect() (err error) {
	once.Do(func() {
		prop := config.Instance.MySQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", prop.Username, prop.Password, prop.Host, prop.Port, prop.Database, prop.Params)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NowFunc: util.Now,
			Logger:  logger.NewGorm(),
		})
	})
	return
}

// AutoMigrate ...
func AutoMigrate(dst ...interface{}) error {
	return DB.AutoMigrate(dst...)
}

// Close ...
func Close() error {
	return nil
}

// Count ...
func Count(table, where string, args ...interface{}) (int, error) {
	var count int
	if err := DB.Raw(fmt.Sprintf("select count(*) from %s where %s", table, where), args...).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Exists ...
func Exists(table, where string, args ...interface{}) bool {
	count, err := Count(table, where, args...)
	if err != nil {
		return false
	}
	return count > 0
}
