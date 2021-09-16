package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"github.com/gofiber/fiber/v2"
)

// CreateUser Create user handler
func CreateUser(c *fiber.Ctx) error {
	var reqUser rr.ReqUser
	if err := c.BodyParser(&reqUser); err != nil {
		return response.MakeFail(c, "param err")
	}

	// param validator
	//isOk := userInfoValidator(reqUser)
	if len(reqUser.Username) <= 0 || len(reqUser.Password) <= 0 || len(reqUser.FirstName) <= 0 || len(reqUser.LastName) <= 0 || len(reqUser.Email) <= 0 || len(reqUser.Phone) <= 0 {
		return response.MakeFail(c, "param err")
	}
	// 用户存在
	flag, _ := redis_read.UserExist(reqUser.Username)
	if flag == 1 {
		return response.MakeFail(c, "user exist")
	}

	flag, err := redis_write.CreateUser(&reqUser)
	if flag != 1 || err != nil {
		return response.MakeFail(c, "insert err")

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

	// 查询用户是否存在,查询用户是否存在并判断密码是否正确
	dbUser, err := redis_read.GetUser(username)
	if err != nil || dbUser == nil || dbUser.Password != password {
		return response.MakeFail(c, "username or password error")
	}

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
