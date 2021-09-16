package redis_write

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"errors"
	"github.com/golang/glog"
	"strings"
)

/* Redis数据结构

所有的房间列表rooms： zset结构  			rooms: [roomID1#roomName1,roomID2#roomName2]
一个房间中的用户列表users：set结构 		roomID: [username1, username2]
用户与房间的对应关系：hashmap结构   room_user : username1: roomID

房间ID与房间名： hashmap结构      room_info: room_id: room_name

*/

var (
	RoomsKey    = "rooms"
	RoomInfoKey = "room_info"
	RoomUserKey = "room_user"
)

// CreateRoom Redis create a room
func CreateRoom(roomID string, roomName string) (int, error) {
	flag, err := rs.ZsPUT(RoomsKey, util.GetSnowflakeInt2(), roomID+"#"+roomName)
	return flag, err
}

// RoomExists room is exist
func RoomExists(roomID string) (int, error) {
	return rs.HExists(RoomInfoKey, roomID)
}

// UserExistRoom user is exists room
func UserExistRoom(userName string) (int, error) {
	return rs.HExists(RoomUserKey, userName)
}

// EnterRoom user enter room
func EnterRoom(roomID string, username string) (int, error) {
	// todo 需要事务控制
	// 放入users列表
	flag, err := rs.SPut(roomID, username)
	if err != nil {
		glog.Error(err)
		return 0, err
	}
	// 放入room_user map
	flag, err = rs.HPut(RoomUserKey, username, roomID)
	return flag, err
}

// LeaveRoom user leave room
func LeaveRoom(roomID string, username string) (int, error) {
	// todo 需要事务控制
	// 清除users列表
	flag, err := rs.SDel(roomID, username)
	if err != nil {
		glog.Error(err)
		return 0, err
	}
	// 清除room_user map
	flag, err = rs.HDel(RoomUserKey, username)
	return flag, err
}

// GetUserInRoom 获取当前用的所在的房间ID
func GetUserInRoom(username string) (string, error) {
	roomID, err := rs.HGet(RoomUserKey, username)
	return roomID.(string), err
}

// GetRoomAllUser 获取此房间所有用户
func GetRoomAllUser(roomID string) ([]string, error) {
	return rs.SGetAll(roomID)
}

// CreateRoomInfo 插入房间信息
func CreateRoomInfo(roomID string, roomName string) (int, error) {
	flag, err := rs.HPut(RoomInfoKey, roomID, roomName)
	return flag, err
}

// GetRoomInfo 获取房间信息
func GetRoomInfo(roomID string) (string, error) {
	roomName, err := rs.HGet(RoomInfoKey, roomID)
	return roomName.(string), err
}

// SelectRoomListPage 分页获取房间列表
func SelectRoomListPage(index, size int) (message []rr.ResRoom, err error) {

	rooms, err := rs.ZsRange(RoomsKey, index, size)
	resRoom := make([]rr.ResRoom, len(rooms))
	for index, value := range rooms {
		values := strings.Split(value, "#")
		resRoom[index].ID = values[0]
		resRoom[index].Name = values[1]
	}
	if len(resRoom) <= 0 || err != nil {
		return nil, errors.New("over max page")
	}
	return resRoom, nil
}
