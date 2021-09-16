package redis_write

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
	"errors"
	"strings"
)

/**
每一个roomID对应一个 ZSet
message信息 ZSet列表  mgs.roomId :[msgId#msgText#msgTime, msgId#msgText#msgTime, ]
*/

func CreateMessage(req *rr.ReqMessage) (int, error) {
	timestamp := util.GetNowTimeUnixNanoString()
	flag, err := rs.ZsPUT("msg."+req.RoomID, util.GetSnowflakeInt2(), req.ID+"##"+req.Text+"##"+timestamp)
	return flag, err
}

func SelectMessageListPage(roomID string, index, size int) (message []rr.ResMessage, err error) {

	msgs, err := rs.ZsRevRange("msg."+roomID, index, size)
	resMessage := make([]rr.ResMessage, len(msgs))
	for index, value := range msgs {
		values := strings.Split(value, "##")
		resMessage[index].ID = values[0]
		resMessage[index].Text = values[1]
		resMessage[index].Timestamp = values[2]
	}
	if len(resMessage) <= 0 || err != nil {
		return nil, errors.New("over max page")
	}
	return resMessage, nil
}
