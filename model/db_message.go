package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"errors"
	"github.com/golang/glog"
)

type Message struct {
	Model
	MessageID string `json:"message_id" gorm:"not null;unique;type:varchar(50)"`
	Text      string `json:"text,omitempty" gorm:"not null;type:varchar(200)"`
	RoomID    string `json:"room_id" gorm:"not null;varchar(50);index:idx_room_id"`
}

// CreateMessage create a message
func CreateMessage(reqMsg *rr.ReqMessage) error {
	var msg = &Message{
		Text: reqMsg.Text,
	}
	msg.MessageID = reqMsg.ID
	msg.RoomID = reqMsg.RoomID

	err := db.Model(&Message{}).Create(msg).Error
	if err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

func SelectMessageListPage(roomID string, index, size int) (message []rr.ResMessage, err error) {

	count := 0
	// 分页子查询
	err = db.Table("message").
		Select("id as idx, message_id as id, text, unix_timestamp(created_at) as timestamp").
		Where("id <= (?)",
			db.Table("message").
				Select("id").
				Where("room_id = ?", roomID).
				Order("id DESC").
				Offset(util.IndexToPage(index, size)).
				Limit(1).
				QueryExpr()).
		Where("room_id = ?", roomID).
		Order("idx DESC").
		Limit(size).
		Count(&count).
		Scan(&message).Error

	// 分页超过范围
	if count < util.IndexToPage(index, size) && len(message) <= 0 {
		return nil, errors.New("over max page")
	}

	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return
}
