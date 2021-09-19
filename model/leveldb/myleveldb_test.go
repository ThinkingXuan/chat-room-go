package myleveldb

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"testing"
)

func init() {
	err := InitLevelDB()
	if err != nil {
		panic(err)
	}
}

func TestCreateUser(t *testing.T) {
	user := &rr.ReqUser{
		Username:  util.GetSnowflakeID2(),
		FirstName: util.GetSnowflakeID2(),
		LastName:  util.GetSnowflakeID2(),
		Email:     util.GetSnowflakeID2(),
		Password:  util.GetSnowflakeID2(),
		Phone:     util.GetSnowflakeID2(),
	}
	err := CreateUser(user)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateRoom(t *testing.T) {

	err := CreateRoom(util.GetSnowflakeID2(), util.GetSnowflakeID2())
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMessage(t *testing.T) {
	user := &rr.ReqMessage{
		ID:     util.GetSnowflakeID2(),
		Text:   util.GetSnowflakeID2(),
		RoomID: util.GetSnowflakeID2(),
	}
	err := CreateMessage(user)
	if err != nil {
		t.Error(err)
	}
}
