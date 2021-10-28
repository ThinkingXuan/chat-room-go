package myleveldb

import (
	"github.com/golang/glog"
)

// 存储方式
// messages.{messagesID}: messageBytes

var (
	messageKey = "messages"
)

// CreateMessage create a message
func CreateMessage(messageID string, reqMsgBytes []byte) error {
	key := messageKey + "." + messageID
	err := db.Put([]byte(key), reqMsgBytes, nil)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

