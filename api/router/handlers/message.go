package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/service"
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

	if err := service.SendMessage(username, &reqMsg); err != nil {
		response.MakeFail(c, err.Error())
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

	messages, err := service.GetMessageList(username, &reqPage)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessJSON(c, messages)
}
