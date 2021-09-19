package myleveldb

import (
	"chat-room-go/api/router/rr"
	"encoding/json"
	"github.com/golang/glog"
)

// 存储方式
// messages.{messagesID}: messageBytes

var (
	messageKey = "messages"
)

// CreateMessage create a message
func CreateMessage(reqMsg *rr.ReqMessage) error {
	key := messageKey + "." + reqMsg.ID
	userBytes, err := json.Marshal(reqMsg)
	err = db.Put([]byte(key), userBytes, nil)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

//func SelectMessageListPage(roomID string, index, size int) (message []rr.ResMessage, err error) {
//
//	count := 0
//	// 分页子查询
//	err = db.Table("message").
//		Select("id as idx, message_id as id, text, unix_timestamp(created_at) as timestamp").
//		Where("id <= (?)",
//			db.Table("message").
//				Select("id").
//				Where("room_id = ?", roomID).
//				Order("id DESC").
//				Offset(util.IndexToPage(index, size)).
//				Limit(1).
//				QueryExpr()).
//		Where("room_id = ?", roomID).
//		Order("idx DESC").
//		Limit(size).
//		Count(&count).
//		Scan(&message).Error
//
//	// 分页超过范围
//	if count < util.IndexToPage(index, size) && len(message) <= 0 {
//		return nil, errors.New("over max page")
//	}
//
//	if err != nil {
//		glog.Error(err)
//		return nil, err
//	}
//	return
//}
