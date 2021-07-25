package middleware

import (
	"chat-room-go/api/router/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuth 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		// 这里的具体实现方式要依据你的实际业务情况决定
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
		//mc, err := jwtauth.ParseToken(token)
		//if err != nil {
		//	response.MakeFail(c, "登录已过期，请重新登录")
		//	c.Abort()
		//	return
		//}
		// 将当前请求的信息保存到请求的上下文gin.context中
		c.Set("username", token)
		c.Next() // 后续的处理函数可以用过c.Get("userAccount")来获取当前请求的用户信息
	}
}
