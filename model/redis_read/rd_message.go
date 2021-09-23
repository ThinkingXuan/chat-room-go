package redis_read

import (
	"chat-room-go/api/router/rr"
	"encoding/json"
	"errors"
)

/**
每一个roomID对应一个 ZSet
message信息 ZSet列表  mgs.roomId :[msgId#msgText#msgTime, msgId#msgText#msgTime, ]
*/

func SelectMessageListPage(roomID string, index, size int) (message []rr.ResMessage, err error) {

	msgs, err := rs.ZsRevRangeBytes("msg."+roomID, index, size)
	var resMessage []rr.ResMessage

	for _, value := range msgs {
		var mesBytes rr.ResMessage
		_ = json.Unmarshal(value, &mesBytes)
		resMessage = append(resMessage, mesBytes)
	}
	if len(resMessage) <= 0 || err != nil {
		return nil, errors.New("over max page")
	}
	return resMessage, nil
}
