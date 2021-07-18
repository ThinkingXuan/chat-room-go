package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model"
	"github.com/gin-gonic/gin"
)

// CreateRoom Create a room
func CreateRoom(c *gin.Context) {
	var reqRoom rr.ReqRoom
	if err := c.ShouldBind(&reqRoom); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}
	if len(reqRoom.Name) <= 0 {
		response.MakeFail(c, "Invalid input")
		return
	}

	room, err := model.CreateRoom(reqRoom.Name)
	if err != nil {
		response.MakeFail(c, "房间创建错误")
		return
	}

	response.MakeSuccess(c, room.ID)
}

// GetOneRoomInfo Get a room information by roomID
func GetOneRoomInfo(c *gin.Context) {

}

// GetRoomList Get room page list
func GetRoomList(c *gin.Context) {

}
