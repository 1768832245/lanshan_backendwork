package flag

import (
	"BM7/lv2/global"
	"BM7/lv2/models"
	"fmt"
)

func DatabaseAutoMigrate() {
	var err error

	//自动建表**
	err = global.DB.Set("gorm:table_option", "Engine=InnoDB").
		AutoMigrate(
			&models.User{},
			&models.Prize{},
			&models.LotteryEntry{},
		)

	if err != nil {
		fmt.Println("自动建表失败")
	} else {
		fmt.Println("自动建表成功")
	}
}
