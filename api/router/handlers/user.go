package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/model"
	"github.com/gin-gonic/gin"
)

// CreateUser Create user handler
func CreateUser(c *gin.Context) {
	var ReqUser rr.ReqUser
	if err := c.ShouldBindJSON(&ReqUser); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if err := model.CreateUser(&ReqUser); err != nil {
		response.MakeFail(c, "添加失败")
		return
	}
	response.MakeSuccessString(c, "注册成功")
}

// UserLogin Logs user into the system handler
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//参数校验
	if len(username) <= 0 || len(password) <= 0 {
		response.MakeFail(c, "Invalid username or password.")
		return
	}

	// 查询用户是否存在
	dbUser, affectRow := model.SelectUserByUsername(username)

	// 查询用户是否存在并判断密码是否正确
	if affectRow <= 0 || password != dbUser.Password {
		response.MakeFail(c, "username or password error")
		return
	}

	tokenString, err := jwtauth.GenToken(dbUser.Username, dbUser.ID)
	if err != nil {
		response.MakeFail(c, "生成Token失败")
		return
	}
	response.MakeSuccessString(c, tokenString)
}

// GetUser Get user by user name handler
func GetUser(c *gin.Context) {
	username := c.Param("username")
	if len(username) <= 0 {
		response.MakeFail(c, "参数错误")
	}
	// 查询用户是否存在
	dbUser, affectRow := model.SelectResUserByUsername(username)
	if affectRow <= 0 {
		response.MakeFail(c, "username error")
		return
	}
	response.MakeSuccessJSON(c, dbUser)
}
