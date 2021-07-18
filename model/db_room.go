package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"github.com/golang/glog"
)

// Room room model
type Room struct {
	Model
	Name string `json:"name,omitempty" gorm:"not null"`
}

// CreateRoom create a room
func CreateRoom(roomName string) (*Room, error) {
	room := &Room{
		Name: roomName,
	}
	room.ID = util.GetSnowflakeID()
	room.CreatedAt = util.GetNowTime()
	err := db.Model(Room{}).Create(room).Error

	if err != nil {
		glog.Error(err)
		return &Room{}, err
	}
	return room, nil
}

func SelectOneRoomByRootName(roomName string) (*rr.ResRoom, int64) {
	var r rr.ResRoom
	rowAffect := db.Table("room").Where("name = ?", roomName).First(&r).RowsAffected
	return &r, rowAffect
}

func SelectRoomListPage(index, size int) (rooms []rr.ResRoom, err error) {
	err = db.Table("room").Order("created_at Desc").Offset(util.IndexToPage(index, size)).Limit(size).
		Scan(rooms).Error
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return
}
