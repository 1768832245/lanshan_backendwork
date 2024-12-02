package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitMysql() *gorm.DB {

	//密码还是不给看哼哼~
	dsn := "root:*******@tcp(127.0.0.1:3306)/BM7?charset=utf8&parseTime=True&loc=Local"

	var mysqlLogger logger.Interface
	mysqlLogger = logger.Default.LogMode(logger.Error)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		fmt.Println("连接失败")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
