package models

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx       = context.Background()
	Stock int = 10
	Rdb   *redis.Client
)
