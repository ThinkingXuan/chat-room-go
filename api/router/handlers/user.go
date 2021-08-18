package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/model/redis"
	"github.com/gin-gonic/gin"
)

// CreateUser Create user handler
func CreateUser(c *gin.Context) {
	var reqUser rr.ReqUser
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	// 用户存在
	flag, _ := redis.UserExist(reqUser.Username)
	if flag == 1 {
		response.MakeFail(c, "user exist")
		return
	}

	flag, err := redis.CreateUser(&reqUser)
	if flag != 1 || err != nil {
		response.MakeFail(c, "insert err")
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

	// 查询用户是否存在,查询用户是否存在并判断密码是否正确
	dbUser, err := redis.GetUser(username)
	if err != nil || dbUser == nil || dbUser.Password != password {
		response.MakeFail(c, "username or password error")
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
	reqUser, err := redis.GetUser(username)
	if err != nil || reqUser == nil {
		response.MakeFail(c, "username error")
		return
	}
	resUser := rr.ResUser{
		FirstName: reqUser.FirstName,
		LastName:  reqUser.LastName,
		Email:     reqUser.Email,
		Phone:     reqUser.Phone,
	}
	response.MakeSuccessJSON(c, resUser)
}
