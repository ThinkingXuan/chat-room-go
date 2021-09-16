package middleware

import (
	"chat-room-go/api/router/response"
	"chat-room-go/internal/jwtauth"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuth 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			response.MakeFail(c, "authorization is null")
			c.Abort()
			return
		}

		// 获取token
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.MakeFail(c, "请求头中auth格式有误")
			c.Abort()
			return
		}

		token := parts[1]
		mc, err := jwtauth.ParseToken(token)
		if err != nil {
			response.MakeFail(c, "登录已过期，请重新登录")
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文gin.context中
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("userAccount")来获取当前请求的用户信息
	}
}
