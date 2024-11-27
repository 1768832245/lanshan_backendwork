package api

import (
	"BM6/lv1Andlv2Andlv3/dao"
	"BM6/lv1Andlv2Andlv3/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	// 重复则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}

func login(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	token, err := utils.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "token generate error",
		})
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
		"token":   token,
	})
}

func GetName(c *gin.Context) {
	// 从上下文中获取用户名
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome!",
		"username": username,
	})
}
