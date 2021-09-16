package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"github.com/gofiber/fiber/v2"
)

// SendMessage user send a message
func SendMessage(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var reqMsg rr.ReqMessage
	if err := c.BodyParser(&reqMsg); err != nil {
		return response.MakeFail(c, "param err")

	}

	if len(reqMsg.ID) <= 0 {
		return response.MakeFail(c, "param err")
	}

	roomID, err := redis_read.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		return response.MakeFail(c, "user no exist room")
	}
	reqMsg.RoomID = roomID

	flag, err := redis_write.CreateMessage(&reqMsg)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "insert err")
	}

	return response.MakeSuccessString(c, "success")

}

// GetMessageList get message list
func GetMessageList(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var reqPage rr.ReqPage
	if err := c.BodyParser(&reqPage); err != nil {
		return response.MakeFail(c, "param err")
	}

	if reqPage.PageIndex >= 0 {
		return response.MakeFail(c, "param err")
	}

	if reqPage.PageSize < 0 {
		return response.MakeFail(c, "param err")
	}

	roomID, err := redis_read.GetUserInRoom(username)
	if err != nil || len(roomID) <= 0 {
		return response.MakeFail(c, "user no exist room")
	}

	messages, err := redis_read.SelectMessageListPage(roomID, reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		return response.MakeFail(c, "select err")
	}
	return response.MakeSuccessJSON(c, messages)
}
