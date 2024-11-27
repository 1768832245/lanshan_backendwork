package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 是一个用于验证 JWT 的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 中获取 Authorization 字段（token）
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 检查 Bearer 格式（我也不知道是啥，似乎就要这么写，在Authorization的[0]是Bearer，[1]是token）
		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		// 解析
		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将解析后的用户信息存入上下文
		c.Set("username", claims.Username)

		// 继续处理请求
		c.Next()
	}
}
