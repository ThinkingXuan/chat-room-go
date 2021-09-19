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

//func SelectOneRoomByRootName(roomName string) (*rr.ResRoom, int64) {
//	var r rr.ResRoom
//	rowAffect := db.Table("room").Where("name = ?", roomName).First(&r).RowsAffected
//	return &r, rowAffect
//}
//
//func SelectOneRoomByRootID(roomID string) (*rr.ResRoom, int64) {
//	var r rr.ResRoom
//	rowAffect := db.Table("room").Where("room_id = ?", roomID).First(&r).RowsAffected
//	return &r, rowAffect
//}
//
//func SelectRoomListPage(index, size int) (rooms []rr.ResRoom, err error) {
//	startIndex := util.IndexToPage(index, size) + 1
//	lastIndex := startIndex + size
//	count := 0
//	err = db.Table("room").
//		Select("room_id as id, name").
//		Order("created_at Desc").
//		Where("id >= ? and id < ?", startIndex, lastIndex).
//		Count(&count).
//		Scan(&rooms).Error
//
//	// 分页超过范围
//	if count < startIndex && len(rooms) <= 0 {
//		return nil, errors.New("over max page")
//	}
//
//	if err != nil {
//		glog.Error(err)
//		return nil, err
//	}
//	return
//}
