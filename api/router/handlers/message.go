package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis"
	"fmt"
	"github.com/gin-gonic/gin"
)

// SendMessage user send a message
func SendMessage(c *gin.Context) {
	username := c.MustGet("username").(string)

	var reqMsg rr.ReqMessage
	if err := c.ShouldBind(&reqMsg); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	if len(reqMsg.ID) <= 0 {
		response.MakeFail(c, "param err")
		return
	}

	roomID, err := redis.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		response.MakeFail(c, "user no exist room")
		return
	}
	reqMsg.RoomID = roomID

	flag, err := redis.CreateMessage(&reqMsg)
	fmt.Println(flag, err)
	if err != nil || flag != 1 {
		response.MakeFail(c, "insert err")
		return
	}

	response.MakeSuccessString(c, "success")

}

// GetMessageList get message list
func GetMessageList(c *gin.Context) {
	username := c.MustGet("username").(string)

	var reqPage rr.ReqPage
	if err := c.ShouldBindJSON(&reqPage); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	if reqPage.PageIndex >= 0 {
		response.MakeFail(c, "param err")
		return
	}

	if reqPage.PageSize < 0 {
		response.MakeFail(c, "param err")
		return
	}

	roomID, err := redis.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		response.MakeFail(c, "user no exist room")
		return
	}

	messages, err := redis.SelectMessageListPage(roomID, reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, "select err")
		return
	}
	response.MakeSuccessJSON(c, messages)
}
