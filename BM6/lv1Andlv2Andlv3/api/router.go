package api

import (
	"BM6/lv1Andlv2Andlv3/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)

	//搜了一下，原来是这样分开使不使用中间件的hhh
	protected := r.Group("/")
	protected.Use(utils.JWTAuthMiddleware())

	protected.GET("/user", GetName)

	errR := r.Run(":8088")
	if errR != nil {
		return
	}
}
