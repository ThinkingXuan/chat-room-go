package middleware

import (
	"chat-room-go/api/router/response"
	"chat-room-go/internal/jwtauth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// JWTAuth 基于JWT的认证中间件
func JWTAuth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return response.MakeFail(c, "authorization is null")
		}

		// 获取token
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return response.MakeFail(c, "请求头中auth格式有误")
		}

		token := parts[1]
		mc, err := jwtauth.ParseToken(token)
		if err != nil {
			return response.MakeFail(c, "登录已过期，请重新登录")
		}
		// 将当前请求的信息保存到请求的上下文context中
		c.Locals("username", mc.Username)
		return c.Next()
	}
}
