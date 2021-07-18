package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"github.com/golang/glog"
)

type Message struct {
	Model
	Text string `json:"text,omitempty" gorm:"not null"`
}

// CreateMessage create a message
func CreateMessage(reqMsg *rr.ReqMessage) error {
	var msg = &Message{
		Text: reqMsg.Text,
	}
	msg.ID = reqMsg.ID
	msg.CreatedAt = util.GetNowTimeUnixNanoString()

	err := db.Model(&Message{}).Create(msg).Error
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func SelectMessageListPage(index, size int) (message []rr.ResMessage, err error) {
	err = db.Table("message").
		Order("created_at Desc").
		Offset(util.IndexToPage(index, size)).
		Limit(size).
		Scan(&message).Error

	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return
}
