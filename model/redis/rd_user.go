package redis

import (
	"chat-room-go/api/router/rr"
	"strings"
)

/*
  用户信息存入redis:  hashmap
  user: username: userInfo
*/

var (
	UserKey = "user"
)

// CreateUser Redis create a user
func CreateUser(reqUser *rr.ReqUser) (int, error) {
	userInfo := reqUser.FirstName + "_#_#" + reqUser.LastName + "_#_#" + reqUser.Email + "_#_#" + reqUser.Password + "_#_#" + reqUser.Phone
	//userInfo := fmt.Sprintf("%s##%s##%s##%s##%s", reqUser.FirstName, reqUs#$#tName, reqUser.Email, reqUser.Password, reqUser.Phone)
	flag, err := rs.HPut(UserKey, reqUser.Username, userInfo)
	return flag, err
}

// UserExist user is exists
func UserExist(userName string) (int, error) {
	return rs.HExists(UserKey, userName)
}

// GetUser get user info form redis
func GetUser(username string) (*rr.ReqUser, error) {
	var reqUser rr.ReqUser
	userInter, err := rs.HGet(UserKey, username)
	if err != nil {
		return &reqUser, err
	}
	userInfo := strings.Split(userInter.(string), "_#_#")
	reqUser.Username = username
	reqUser.FirstName = userInfo[0]
	reqUser.LastName = userInfo[1]
	reqUser.Email = userInfo[2]
	reqUser.Password = userInfo[3]
	reqUser.Phone = userInfo[4]

	return &reqUser, nil
}
