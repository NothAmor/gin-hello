package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/NothAmor/gin-hello/app/models"
)

func InitDatabase() *gorm.DB {
	// 配置数据库连接dsn
	dsn := "gin-hello:wScRz3KRh2xxbRXR@tcp(47.101.38.190:3306)/gin-hello?charset=utf8&parseTime=True&loc=Local"

	// 建立数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func DatabaseMigrate(db *gorm.DB, migrate bool) {
	if migrate {
		db.AutoMigrate(&models.Users{})
	}
}
