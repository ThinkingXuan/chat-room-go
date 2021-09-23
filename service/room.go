package service

import (
	"chat-room-go/api/router/rr"
	myleveldb "chat-room-go/model/leveldb"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"encoding/json"
	"errors"
)

func CreateRoom(roomID string, roomName string) error {

	// 序列化
	resRoom := rr.ResRoom{
		Name: roomName,
		ID:   roomID,
	}
	resRoomByte, _ := json.Marshal(&resRoom)

	// room写入leveldb
	err := myleveldb.CreateRoom(roomID, roomName)
	if err != nil {
		return err
	}

	// room写入redis
	// write to redis： crate a room zset
	//flag, err := redis_write.CreateRoom(resRoomByte)
	//if err != nil || flag != 1 {
	//	return errors.New("create room err")
	//}
	////	write to redis： crate a room info
	//flag, err = redis_write.CreateRoomInfo(roomID, roomName)
	//if err != nil || flag != 1 {
	//	return errors.New("create room err")
	//}
	err = redis_write.CreateRoomAndRoomInfo(roomID, roomName, resRoomByte)
	if err != nil {
		return err
	}
	return nil
}

func GetOneRoomInfo(roomID string) (string, error) {
	return redis_read.GetRoomInfo(roomID)
}

// GetRoomList Get room page list
func GetRoomList(index, size int) ([]rr.ResRoom, error) {

	rooms, err := redis_read.SelectRoomListPage(index, size)
	if err != nil {
		return nil, errors.New("select err")
	}
	return rooms, nil
}

// EnterRoom user enter room
func EnterRoom(roomID string, username string) error {

	// 判断房间是否存在
	flag, err := redis_read.RoomExists(roomID)
	if err != nil || flag != 1 {
		return errors.New("invalid room ID")
	}

	//获取所在房间的ID
	oldRoomID, _ := redis_read.GetUserInRoom(username)

	// 现在所在房间为要进入的房间
	if oldRoomID == roomID {
		return nil
	}
	// 现在所在房间为不是自己要进入的房间
	if len(oldRoomID) > 0 {
		// 离开房间
		err := redis_write.LeaveRoomMerge(oldRoomID, username)
		if err != nil {
			return errors.New("leave Room failure")
		}
	}
	// 进入此房间
	err = redis_write.EnterRoomMerge(roomID, username)
	if err != nil {
		return errors.New("enter room failure")

	}
	return nil
}

// LeaveRoom user leaven room
func LeaveRoom(username string) error {
	//获取所在房间的ID
	oldRoomID, _ := redis_read.GetUserInRoom(username)
	// 用户不在房间
	if len(oldRoomID) <= 0 {
		return errors.New("leave Room failure")
	}
	err := redis_write.LeaveRoomMerge(oldRoomID, username)
	if err != nil {
		return errors.New("leave Room failure")
	}
	return nil
}

// RoomAllUser get room all user list
func RoomAllUser(roomID string) ([]string, error) {
	roomUser, err := redis_read.GetRoomAllUser(roomID)
	if err != nil || len(roomUser) < 0 {
		return nil, errors.New("get users failure")
	}
	return roomUser, nil
}
