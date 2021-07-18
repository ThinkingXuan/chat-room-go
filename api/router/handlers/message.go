package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model"
	"github.com/gin-gonic/gin"
)

// SendMessage user send a message
func SendMessage(c *gin.Context) {
	var reqMsg rr.ReqMessage
	if err := c.ShouldBind(&reqMsg); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if len(reqMsg.ID) <= 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	err := model.CreateMessage(&reqMsg)
	if err != nil {
		response.MakeFail(c, "插入错误")
		return
	}

	response.MakeSuccess(c, "success")
}

// GetMessageList get message list
func GetMessageList(c *gin.Context) {
	var reqPage rr.ReqPage
	if err := c.ShouldBindJSON(&reqPage); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if reqPage.PageSize < 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	messages, err := model.SelectMessageListPage(reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, "查询错误")
		return
	}
	response.MakeSuccess(c, messages)
}
