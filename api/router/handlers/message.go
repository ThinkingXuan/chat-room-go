package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model"
	"chat-room-go/model/redis"
	"github.com/gin-gonic/gin"
)

// SendMessage user send a message
func SendMessage(c *gin.Context) {
	username := c.MustGet("username").(string)

	var reqMsg rr.ReqMessage
	if err := c.ShouldBind(&reqMsg); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if len(reqMsg.ID) <= 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	roomID, err := redis.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		response.MakeFail(c, "user no exist room")
		return
	}
	reqMsg.RoomID = roomID

	err = model.CreateMessage(&reqMsg)
	if err != nil {
		response.MakeFail(c, "插入错误")
		return
	}

	response.MakeSuccessString(c, "success")
}

// GetMessageList get message list
func GetMessageList(c *gin.Context) {
	username := c.MustGet("username").(string)

	var reqPage rr.ReqPage
	if err := c.ShouldBindJSON(&reqPage); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if reqPage.PageIndex >= 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	if reqPage.PageSize < 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	roomID, err := redis.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		response.MakeFail(c, "user no exist room")
		return
	}

	messages, err := model.SelectMessageListPage(roomID, reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, "查询错误")
		return
	}
	response.MakeSuccessJSON(c, messages)
}
