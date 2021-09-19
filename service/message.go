package service

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"errors"
)

// SendMessage user send a message
func SendMessage(username string, reqMsg *rr.ReqMessage) error {

	roomID, err := redis_read.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		return errors.New("user no exist room")
	}
	reqMsg.RoomID = roomID

	// message写入leveldb
	//err = myleveldb.CreateMessage(reqMsg)
	//if err != nil {
	//	return errors.New("leveldb insert err")
	//}
	// message写入redis
	flag, err := redis_write.CreateMessage(reqMsg)
	if err != nil || flag != 1 {
		return errors.New("insert err")
	}

	return nil
}

// GetMessageList get message list
func GetMessageList(username string, reqPage *rr.ReqPage) ([]rr.ResMessage, error) {

	roomID, err := redis_read.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		return nil, errors.New("user no exist room")

	}
	messages, err := redis_read.SelectMessageListPage(roomID, reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		return nil, errors.New("select err")
	}
	return messages, nil
}
