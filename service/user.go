package service

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"encoding/json"
	"errors"
)

func CreateUser(reqUser *rr.ReqUser) error {
	// 用户存在时CreateUser时 flag = 0
	//flag, _ := redis_read.UserExist(reqUser.Username)
	//if flag == 1 {
	//	return errors.New("user exist")
	//}

	// 写入leveldb
	reqUserBytes, _ := json.Marshal(reqUser)
	//err := myleveldb.CreateUser(reqUser.Username, reqUserBytes)
	//if err != nil {
	//	return errors.New("leveldb insert err")
	//}
	// 写入redis
	flag, err := redis_write.CreateUser(reqUser.Username, reqUserBytes)
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

func GetUser(username string) (*rr.ResUser, error) {
	reqUser, err := redis_read.GetUser(username)
	if err != nil || reqUser == nil {
		return nil, errors.New("username error")
	}
	resUser := &rr.ResUser{
		FirstName: reqUser.FirstName,
		LastName:  reqUser.LastName,
		Email:     reqUser.Email,
		Phone:     reqUser.Phone,
	}
	return resUser, nil
}
