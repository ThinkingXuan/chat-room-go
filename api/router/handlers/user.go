package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/service"
	"github.com/gin-gonic/gin"
)

// CreateUser Create user handler
func CreateUser(c *gin.Context) {
	var reqUser rr.ReqUser
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	// param validator
	if len(reqUser.Username) <= 0 || len(reqUser.Password) <= 0 || len(reqUser.FirstName) <= 0 || len(reqUser.LastName) <= 0 || len(reqUser.Email) <= 0 || len(reqUser.Phone) <= 0 {
		response.MakeFail(c, "param err")
		return
	}

	// 创建用户
	err := service.CreateUser(&reqUser)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessString(c, "successful operation")
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

	err := service.UserLogin(username, password)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}

	tokenString, err := jwtauth.GenToken(username)
	if err != nil {
		response.MakeFail(c, "generate jwt token failed")
		return
	}
	response.MakeSuccessString(c, tokenString)
}

// GetUser Get user by user name handler
func GetUser(c *gin.Context) {
	username := c.Param("username")
	if len(username) <= 0 {
		response.MakeFail(c, "param err")
		return
	}
	// 查询用户是否存在

	reqUser, err := service.GetUser(username)
	if err != nil || reqUser == nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessJSON(c, reqUser)
}
