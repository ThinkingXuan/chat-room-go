package redis_write

import (
	"chat-room-go/util"
	"github.com/golang/glog"
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

// CreateRoomAndRoomInfo Redis create a room
func CreateRoomAndRoomInfo(roomID string, roomName string) error {
	err := rs.CreateRoomAndRoomInfo(roomID, roomName)
	return err
}

// CreateRoom Redis create a room
func CreateRoom(roomID string, roomName string) (int, error) {
	flag, err := rs.ZsPUT(RoomsKey, util.GetSnowflakeInt2(), roomID+"#"+roomName)
	return flag, err
}

// CreateRoomInfo 插入房间信息
func CreateRoomInfo(roomID string, roomName string) (int, error) {
	flag, err := rs.HPut(RoomInfoKey, roomID, roomName)
	return flag, err
}

// EnterRoom user enter room
func EnterRoom(roomID string, username string) (int, error) {
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
