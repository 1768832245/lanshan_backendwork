package main

import (
	"BM8/lv4/api"
	"BM8/lv4/dao"
	"BM8/lv4/models"
	"bufio"
	"os"
)

func main() {
	dao.InitRdb()

	go api.Sub(models.Ctx, "kq")

	for i := 0; i < 5; i++ {
		msg := bufio.NewReader(os.Stdin)
		input, _ := msg.ReadString('\n')
		api.Pub(models.Ctx, "kq", input)
	}
}
