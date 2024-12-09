package main

import (
	"BM8/lv3/api"
	"BM8/lv3/dao"
	"time"
)

func main() {
	dao.InitRdb()

	for i := 1; i <= 5; i++ {
		go api.PurchaseStock(i)
	}

	time.Sleep(5 * time.Second)
}
