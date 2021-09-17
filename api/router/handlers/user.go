package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/model/redis_read"
	"chat-room-go/service"
	"github.com/gofiber/fiber/v2"
)

// CreateUser Create user handler
func CreateUser(c *fiber.Ctx) error {
	var reqUser rr.ReqUser
	if err := c.BodyParser(&reqUser); err != nil {
		return response.MakeFail(c, "param err")
	}

	// param validator
	if len(reqUser.Username) <= 0 || len(reqUser.Password) <= 0 || len(reqUser.FirstName) <= 0 || len(reqUser.LastName) <= 0 || len(reqUser.Email) <= 0 || len(reqUser.Phone) <= 0 {
		return response.MakeFail(c, "param err")
	}

	err := service.CreateUser(&reqUser)
	if err != nil {
		return response.MakeFail(c, err.Error())
	}

	return response.MakeSuccessString(c, "successful operation")
}

// UserLogin Logs user into the system handler
func UserLogin(c *fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")

	//参数校验
	if len(username) <= 0 || len(password) <= 0 {
		return response.MakeFail(c, "Invalid username or password.")
	}
	err := service.UserLogin(username, password)
	if err != nil {
		return response.MakeFail(c, err.Error())
	}
	// JWT Token
	tokenString, err := jwtauth.GenToken(username)
	if err != nil {
		return response.MakeFail(c, "generate jwt token failed")
	}
	return response.MakeSuccessString(c, tokenString)
}

// GetUser Get user by user name handler
func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	if len(username) <= 0 {
		return response.MakeFail(c, "param err")
	}
	// 查询用户是否存在
	reqUser, err := redis_read.GetUser(username)
	if err != nil || reqUser == nil {
		return response.MakeFail(c, "username error")
	}
	resUser := rr.ResUser{
		FirstName: reqUser.FirstName,
		LastName:  reqUser.LastName,
		Email:     reqUser.Email,
		Phone:     reqUser.Phone,
	}
	return response.MakeSuccessJSON(c, resUser)
}
