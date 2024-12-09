package main

import (
	"BM8/lv2/api"
	"BM8/lv2/dao"
	"fmt"
)

func main() {
	dao.InitDb()
	
	user, err := api.GetUserInfo(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*user)
}
