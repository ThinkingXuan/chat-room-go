package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"errors"
	"github.com/golang/glog"
)

// Room room model
type Room struct {
	Model
	RoomID string `json:"room_id" gorm:"not null;type:varchar(50)"`
	Name   string `json:"name,omitempty" gorm:"not null;type:varchar(50)"`
}

// CreateRoom create a room
func CreateRoom(roomID string, roomName string) (*Room, error) {
	room := &Room{
		Name: roomName,
	}
	room.RoomID = roomID
	err := db.Model(&Room{}).Create(room).Error
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return room, nil
}

// CreateAsyncRoom create a room
func CreateAsyncRoom(roomID string, roomName string) {
	room := &Room{
		Name: roomName,
	}
	room.RoomID = roomID
	db.Model(&Room{}).Create(room)
	//if err != nil {
	//	glog.Error(err)
	//	return nil, err
	//}
	//return room, nil
}

func SelectOneRoomByRootName(roomName string) (*rr.ResRoom, int64) {
	var r rr.ResRoom
	rowAffect := db.Table("room").Where("name = ?", roomName).First(&r).RowsAffected
	return &r, rowAffect
}

func SelectOneRoomByRootID(roomID string) (*rr.ResRoom, int64) {
	var r rr.ResRoom
	rowAffect := db.Table("room").Where("room_id = ?", roomID).First(&r).RowsAffected
	return &r, rowAffect
}

func SelectRoomListPage(index, size int) (rooms []rr.ResRoom, err error) {
	startIndex := util.IndexToPage(index, size) + 1
	lastIndex := startIndex + size
	count := 0
	err = db.Table("room").
		Select("room_id as id, name").
		Order("created_at Desc").
		Where("id >= ? and id < ?", startIndex, lastIndex).
		Count(&count).
		Scan(&rooms).Error

	// 分页超过范围
	if count < startIndex && len(rooms) <= 0 {
		return nil, errors.New("over max page")
	}

	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return
}
