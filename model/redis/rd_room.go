package redis

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var (
	rsRoomKey     = viper.GetString("redis.room_key")
	rsRoomUserKey = viper.GetString("redis.room_user_key")
)

/* Redis数据结构

所有的房间列表rooms： set结构  			rooms: [roomID1,roomID2]
一个房间中的用户列表users：set结构 		roomID: [username1, username2]
用户与房间的对应关系：hashmap结构   room_user : username1: roomID

*/

// CreateRoom Redis create a room
func CreateRoom(roomID string, roomName string) (int, error) {
	return rs.SPut(rsRoomKey, roomID)
}

// RoomExists room is exist
func RoomExists(roomID string) (int, error) {
	return rs.SExists(rsRoomKey, roomID)
}

// UserExistRoom user is exists room
func UserExistRoom(userName string) (int, error) {
	return rs.HExists(rsRoomUserKey, userName)
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
	flag, err = rs.HPut(rsRoomUserKey, username, roomID)
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
	flag, err = rs.HDel(rsRoomUserKey, username)
	return flag, err
}

// GetUserInRoom 获取当前用的所在的房间ID
func GetUserInRoom(username string) (string, error) {
	roomID, err := rs.HGet(rsRoomUserKey, username)
	return roomID.(string), err
}

// GetRoomAllUser 获取此房间所有用户
func GetRoomAllUser(roomID string) ([]string, error) {
	return rs.SGetAll(roomID)
}
