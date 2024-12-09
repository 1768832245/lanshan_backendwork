package dao

import (
	"BM8/lv2/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDb() {
	var err0 error
	dsn := "root:********@tcp(localhost:3306)/bm8?charset=utf8mb4&parseTime=True&loc=Local"
	models.Db, err0 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err0 != nil {
		fmt.Println("mysql init err:", err0)
	}

	models.Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "",
		DB:       0})

	_, err := models.Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("redis 连接成功")
}
