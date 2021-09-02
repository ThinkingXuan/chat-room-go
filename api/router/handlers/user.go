package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/jwtauth"
	"chat-room-go/model/redis"
	"chat-room-go/model/redis_read"
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
	//isOk := userInfoValidator(reqUser)
	if len(reqUser.Username) <= 0 || len(reqUser.Password) <= 0 || len(reqUser.FirstName) <= 0 || len(reqUser.LastName) <= 0 || len(reqUser.Email) <= 0 || len(reqUser.Phone) <= 0 {
		response.MakeFail(c, "param err")
		return
	}
	// 用户存在
	flag, _ := redis_read.UserExist(reqUser.Username)
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

//var (
//	regEmail     = regexp2.MustCompile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", regexp2.RE2)
//	regPhone     = regexp2.MustCompile("^1(3[0-9]|5[0-3,5-9]|7[1-3,5-8]|8[0-9])\\d{8}$", regexp2.RE2)
//	regUserName  = regexp2.MustCompile("^[\\\\u4e00-\\\\u9fa5_a-zA-Z0-9-]{1,16}$", regexp2.RE2)
//	regPassword  = regexp2.MustCompile("^(?![a-zA-Z]+$)(?!\\d+$)(?![!@#$%^&*]+$)[a-zA-Z\\d!@#()_$%-^.&*]{6,20}$", regexp2.RE2)
//	regFirstName = regexp2.MustCompile("^[\\u4e00-\\u9fa5_a-zA-Z]+$", regexp2.RE2)
//	regLastName  = regexp2.MustCompile("^[\\u4e00-\\u9fa5_a-zA-Z]+$", regexp2.RE2)
//)
//
//func userInfoValidator(user rr.ReqUser) bool {
//	// email
//	emailOk, _ := regEmail.MatchString(user.Email)
//	// phone
//	phoneOk, _ := regPhone.MatchString(user.Phone)
//	// username
//	usernameOK, _ := regUserName.MatchString(user.Username)
//	passwordOK, _ := regPassword.MatchString(user.Password)
//
//	firstNameOk, _ := regFirstName.MatchString(user.FirstName)
//	lastNameOk, _ := regLastName.MatchString(user.LastName)
//
//	return emailOk && phoneOk && usernameOK && firstNameOk && lastNameOk && passwordOK
//}

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
	dbUser, err := redis_read.GetUser(username)
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
	reqUser, err := redis_read.GetUser(username)
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
