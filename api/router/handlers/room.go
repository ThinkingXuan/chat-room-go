package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model"
	"fmt"
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
	roomID := c.Param("roomid")
	if len(roomID) <= 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	resRoom, affectRow := model.SelectOneRoomByRootID(roomID)
	fmt.Println(resRoom)
	if affectRow <= 0 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}
	response.MakeSuccess(c, resRoom.Name)
}

// GetRoomList Get room page list
func GetRoomList(c *gin.Context) {
	var reqPage rr.ReqPage
	if err := c.ShouldBindJSON(&reqPage); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}

	if reqPage.PageSize < 0 || reqPage.PageIndex < 0 {
		response.MakeFail(c, "参数错误")
		return
	}

	rooms, err := model.SelectRoomListPage(reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, "查询错误")
		return
	}
	response.MakeSuccess(c, rooms)
}
