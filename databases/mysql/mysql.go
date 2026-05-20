// MySQL 数据库连接管理包。
//
// 通过 init() 自动连接数据库、配置连接池、执行自动迁移。
// 连接成功后可通过 Db 全局变量使用 GORM 实例。
package mysql

import (
	"app/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// Db 是全局 GORM 数据库实例，由 init() 初始化。
var Db *gorm.DB

// 数据库连接配置，由环境变量（.env）注入。
var (
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbDatabase string
)

func init() {
	settings.RequireEnvs([]string{
		"DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD",
	})

	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")

	var err error

	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUsername,
		DbPassword,
		DbHost,
		DbPort,
		DbDatabase,
	)

	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(55 * time.Second)

	if settings.Debug {
		Db = Db.Debug()
	}

	err = Db.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
