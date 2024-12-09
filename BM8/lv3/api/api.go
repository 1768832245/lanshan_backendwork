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
		fmt.Println("已被占据")
		return false
	}
	if models.Stock > 0 {
		models.Stock--
		fmt.Println("success stock:", models.Stock)
	} else {
		fmt.Println("stock < 0!")
	}

	models.Rdb.Del(models.Ctx, lock)
	fmt.Println("已完成释放锁")
	return true
}
