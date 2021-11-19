package simple_factory

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MysqlDriver = "mysql"
	PostgresDriver = "postgres"
)

var (
	gormDB *gorm.DB
)

type DBSelect interface {
	DBSelect()
}

type Database struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBDatabase string
}

func mysqlSelect() (gorm.Dialector, error) {
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUser, DBPassword,
		DBHost, DBPort, DBDatabase)
	if MysqlDriverParams != "" {
		address = fmt.Sprintf("%s?%s", address, MysqlDriverParams)
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       address, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	})

	return dialector, nil
}

func postgresSelect() (gorm.Dialector, error) {
	address := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable client_encoding=UTF8 TimeZone=Asia/Shanghai", DBUser,
		DBPassword, DBDatabase, DBHost, DBPort)

	dialector := postgres.New(postgres.Config{
		DSN:                  address,
		PreferSimpleProtocol: true,
	})

	return dialector, nil
}

func (d *Database) DBSelect() {
	var dialector gorm.Dialector
	switch DBDriver {
	case MysqlDriver:
		dialector, _ = mysqlSelect()
	case PostgresDriver:
		dialector, _ = postgresSelect()
	}
	gormDB, _ = gorm.Open(dialector)
}

func DB() *gorm.DB {
	if DBDebugMode {
		return gormDB.Debug()
	}
	return gormDB
}

func GetDB() *gorm.DB {
	// return gormDB.Debug()
	return DB()
}
