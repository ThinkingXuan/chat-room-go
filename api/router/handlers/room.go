package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/service"
	"chat-room-go/util"
	"github.com/gin-gonic/gin"
)

// CreateRoom Create a room
func CreateRoom(c *gin.Context) {
	var reqRoom rr.ReqRoom
	if err := c.ShouldBindJSON(&reqRoom); err != nil {
		response.MakeFail(c, "param err")
		return
	}
	if len(reqRoom.Name) <= 0 {
		response.MakeFail(c, "Invalid input")
		return
	}

	// create a room id
	roomID := util.GetSnowflakeID2()
	//	write to rate a room info
	err := service.CreateRoom(roomID, reqRoom.Name)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessString(c, roomID)

}

// GetOneRoomInfo Get a room information by roomID
func GetOneRoomInfo(c *gin.Context) {
	roomID := c.Param("roomid")
	if len(roomID) <= 0 {
		response.MakeFail(c, "param err")
		return
	}

	roomName, err := service.GetOneRoomInfo(roomID)
	if len(roomName) <= 0 || err != nil {
		response.MakeFail(c, "Invalid Room ID")
		return
	}
	response.MakeSuccessString(c, roomName)
}

// GetRoomList Get room page list
func GetRoomList(c *gin.Context) {
	var reqPage rr.ReqPage
	if err := c.ShouldBindJSON(&reqPage); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	if reqPage.PageSize < 0 || reqPage.PageIndex < 0 {
		response.MakeFail(c, "param err")
		return
	}

	rooms, err := service.GetRoomList(reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessJSON(c, rooms)
}

// EnterRoom user enter room
func EnterRoom(c *gin.Context) {
	username := c.MustGet("username").(string)
	roomID := c.Param("roomid")
	// 参数校验
	if len(roomID) <= 0 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}

	err := service.EnterRoom(roomID, username)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}

	response.MakeSuccessString(c, "enter the Room success")
}

// LeaveRoom user leaven room
func LeaveRoom(c *gin.Context) {
	username := c.MustGet("username").(string)

	err := service.LeaveRoom(username)
	if err != nil {
		response.MakeFail(c, err.Error())
		return
	}
	response.MakeSuccessString(c, "left the room")
}

// RoomAllUser get room all user list
func RoomAllUser(c *gin.Context) {
	roomID := c.Param("roomid")
	if len(roomID) <= 0 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}
	roomUser, err := service.RoomAllUser(roomID)
	if err != nil || len(roomUser) < 0 {
		response.MakeFail(c, "get users failure")
		return
	}
	response.MakeSuccessJSON(c, roomUser)
}
