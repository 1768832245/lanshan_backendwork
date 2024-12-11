package api

import (
	"BM8/lv3/models"
	"fmt"
	"time"
)

// 模拟秒杀
func PurchaseStock(userid int) bool {
	lock := "lock"
	LS, err := models.Rdb.SetNX(models.Ctx, lock, userid, 100*time.Second).Result()
	if err != nil {
		return false
	}

	if !LS {
		return false
	}
	if models.Stock > 0 {
		models.Stock--
		fmt.Println("success stock:", models.Stock)
		fmt.Printf("userid:%d,抢到了货物\n", userid)
	} else {
		models.Rdb.Del(models.Ctx, lock)
		return false
	}

	models.Rdb.Del(models.Ctx, lock)
	return true
}
