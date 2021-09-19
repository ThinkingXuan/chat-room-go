package redis_write

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
func CreateUser(username string, newUserBytes []byte) (int, error) {
	flag, err := rs.HPut(UserKey, username, newUserBytes)
	return flag, err
}

// UserExist user is exists
func UserExist(userName string) (int, error) {
	return rs.HExists(UserKey, userName)
}

// GetUser get user info form redis_write
func GetUser(username string) (*rr.ReqUser, error) {
	var reqUser rr.ReqUser
	userInter, err := rs.HGet(UserKey, username)
	if err != nil {
		return &reqUser, err
	}
	userInfo := strings.Split(userInter.(string), "##")
	reqUser.Username = username
	reqUser.FirstName = userInfo[0]
	reqUser.LastName = userInfo[1]
	reqUser.Email = userInfo[2]
	reqUser.Password = userInfo[3]
	reqUser.Phone = userInfo[4]

	return &reqUser, nil
}
