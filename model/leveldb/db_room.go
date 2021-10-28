package myleveldb

import (
	"github.com/golang/glog"
)

// 存储方式
// rooms.{roomID}: roomName

var (
	roomKeys = "rooms"
)

// CreateRoom create a room
func CreateRoom(roomID string, roomName string) error {
	key := roomKeys + "." + roomID
	err := db.Put([]byte(key), []byte(roomName), nil)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}
