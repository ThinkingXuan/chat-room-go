package redis_read

import (
	"chat-room-go/api/router/rr"
	"encoding/json"
)

/*
  用户信息存入redis:  hashmap
  user: username: userInfo
*/

var (
	UserKey = "user"
)

// GetUser get user info form redis_write
func GetUser(username string) (*rr.ReqUser, error) {
	var reqUser rr.ReqUser
	userBytes, err := rs.HGetBytes(UserKey, username)
	if err != nil {
		return &reqUser, err
	}

	// 反序列化
	err = json.Unmarshal(userBytes, &reqUser)
	if err != nil {
		return &reqUser, err
	}

	return &reqUser, nil
}
