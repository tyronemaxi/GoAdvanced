package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 连接数据库
// 定义数据库全局
var once sync.Once

var db *database

type database struct {
	instance    *gorm.DB
	maxIdle     int
	maxOpen     int
	maxLifetime time.Duration
}
type Option func(db *database)

func DB(opts ...Option) *gorm.DB {
	once.Do(func() {
		db = new(database)
		for _, f := range opts {
			f(db)
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost", 3306, "tyrone")

		var err error
		db.instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		//
		//sqlDB,err := db.instance.DB()
		//if err != nil {
		//	panic(err)
		//}

	})
	return db.instance
}
