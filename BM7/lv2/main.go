// main.go
package main

import (
	"BM7/lv2/core"
	"BM7/lv2/flag"
	"BM7/lv2/global"
	"BM7/lv2/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// 初始化数据库
	global.DB = core.InitMysql()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 创建 Gin 实例
	r := gin.Default()

	// 查看所有奖品
	r.GET("/prizes", func(c *gin.Context) {
		var prizes []models.Prize
		global.DB.Find(&prizes)
		c.JSON(http.StatusOK, prizes)
	})

	// 创建用户
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		global.DB.Create(&user)
		c.JSON(http.StatusCreated, user)
	})

	// 创建奖品
	r.POST("/prizes", func(c *gin.Context) {
		var prize models.Prize
		if err := c.ShouldBindJSON(&prize); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		global.DB.Create(&prize)
		c.JSON(http.StatusCreated, prize)
	})

	r.POST("/lottery/enter", func(c *gin.Context) {
		var prizes []models.Prize
		var request struct {
			UserID uint `json:"userid"` // 用户ID
		}

		// 获取UserID
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		//奖品全部传入prizes数组
		global.DB.Where("stock > 0").Find(&prizes)

		// 创建随机数生成器，基于当前时间生成种子
		source := rand.NewSource(time.Now().UnixNano())
		randGen := rand.New(source)

		// 根据随机数种子随机选择一个奖品（搜的hhhh）
		prize := prizes[randGen.Intn(len(prizes))]

		// 顺便创建抽奖记录(
		lotteryEntry := models.LotteryEntry{
			UserID:    request.UserID,
			PrizeID:   prize.ID,
			Timestamp: time.Now().Unix(),
		}
		global.DB.Create(&lotteryEntry)

		// 更新库存
		prize.Stock--

		//根据主键自动存储
		global.DB.Save(&prize)

		c.JSON(200, gin.H{
			"message":     "successfully",
			"lottery_id":  lotteryEntry.ID,
			"prize_name":  prize.Name,
			"prize_stock": prize.Stock,
		})
	})

	//显示抽奖记录
	r.GET("/lottery", func(c *gin.Context) {
		var lotteryEntries []models.LotteryEntry
		//user载体
		var user models.User
		//prize载体
		var prize models.Prize
		global.DB.Find(&lotteryEntries)
		for _, lotteryEntry := range lotteryEntries {

			//用id查询
			global.DB.Find(&prize, lotteryEntry.PrizeID)
			global.DB.Find(&user, lotteryEntry.UserID)

			c.JSON(200, gin.H{
				"message":    "successfully",
				"lottery_id": lotteryEntry.ID,
				"userid":     lotteryEntry.UserID,
				"prize_id":   lotteryEntry.PrizeID,
				"timestamp":  lotteryEntry.Timestamp,
				"user":       user,
				"prize":      prize,
			})
		}
	})

	R := r.Run(":8080")
	if R != nil {
		return
	}
}
