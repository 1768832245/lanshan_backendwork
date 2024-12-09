package models

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
	Db  *gorm.DB
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"` // 自增ID
	Name  string `gorm:"type:varchar(255);" json:"name"`
	Email string `gorm:"type:varchar(255);" json:"email"`
}
