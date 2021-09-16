package redis_read

import (
	"chat-room-go/api/router/rr"
	"errors"
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

// RoomExists room is exist
func RoomExists(roomID string) (int, error) {
	//rsRoomKey := viper.GetString("redis_write.room_key")
	return rs.HExists(RoomInfoKey, roomID)
	//return rs.SExists(rsRoomKey, roomID)
}

// UserExistRoom user is exists room
func UserExistRoom(userName string) (int, error) {

	return rs.HExists(RoomUserKey, userName)
}

// GetUserInRoom 获取当前用户所在的房间ID
func GetUserInRoom(username string) (string, error) {
	roomID, err := rs.HGet(RoomUserKey, username)
	return roomID.(string), err
}

// GetRoomAllUser 获取此房间所有用户
func GetRoomAllUser(roomID string) ([]string, error) {
	return rs.SGetAll(roomID)
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
