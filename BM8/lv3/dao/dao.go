package dao

import (
	"BM8/lv3/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func InitRdb() {
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
