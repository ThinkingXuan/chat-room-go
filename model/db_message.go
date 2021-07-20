package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"github.com/golang/glog"
)

type Message struct {
	Model
	Text   string `json:"text,omitempty" gorm:"not null"`
	RoomID string `json:"room_id" gorm:"not null"`
}

// CreateMessage create a message
func CreateMessage(reqMsg *rr.ReqMessage) error {
	var msg = &Message{
		Text: reqMsg.Text,
	}
	msg.ID = reqMsg.ID
	msg.CreatedAt = util.GetNowTimeUnixNanoString()
	msg.RoomID = reqMsg.RoomID

	err := db.Model(&Message{}).Create(msg).Error
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func SelectMessageListPage(roomID string, index, size int) (message []rr.ResMessage, err error) {
	err = db.Table("message").
		Select("id, text, created_at as timestamp").
		Where("room_id = ?", roomID).
		Order("timestamp Desc").
		Offset(util.IndexToPage(index, size)).
		Limit(size).
		Scan(&message).Error

	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return
}
