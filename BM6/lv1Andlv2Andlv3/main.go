package main

import (
	"BM6/lv1Andlv2Andlv3/api"
	"BM6/lv1Andlv2Andlv3/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
