package global

import (
	"BM7/lv2/conf"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
