package service

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"errors"
)

func CreateUser(reqUser *rr.ReqUser) error {
	// 用户存在
	flag, _ := redis_read.UserExist(reqUser.Username)
	if flag == 1 {
		return errors.New("user exist")
	}

	flag, err := redis_write.CreateUser(reqUser)
	if flag != 1 || err != nil {
		return errors.New("insert err")

	}
	return nil
}

func UserLogin(username string, password string) error {
	// 查询用户是否存在,查询用户是否存在并判断密码是否正确
	dbUser, err := redis_read.GetUser(username)
	if err != nil || dbUser == nil || dbUser.Password != password {
		return errors.New("username or password error")
	}
	return nil
}
