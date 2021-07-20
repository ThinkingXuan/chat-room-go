package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model"
	"chat-room-go/model/redis"
	"chat-room-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateRoom Create a room
func CreateRoom(c *gin.Context) {
	//username := c.MustGet("username").(string)
	var reqRoom rr.ReqRoom
	if err := c.ShouldBind(&reqRoom); err != nil {
		response.MakeFail(c, "参数错误")
		return
	}
	if len(reqRoom.Name) <= 0 {
		response.MakeFail(c, "Invalid input")
		return
	}

	//// 判断此用户是否已在房间中
	//flag, _ := redis.UserExistRoom(username)
	//if flag == 1 { // 在房间中,离开
	//	//获取所在房间的ID
	//	oldRoomID, _ := redis.GetUserInRoom(username)
	//	_, err := redis.LeaveRoom(oldRoomID, username)
	//	if err != nil {
	//		response.MakeFail(c, "leave Room failure")
	//	}
	//}

	// create a room id
	roomID := util.GetSnowflakeID()

	// write to redis
	flag, err := redis.CreateRoom(roomID)
	if err != nil || flag != 1 {
		response.MakeFail(c, "房间创建错误")
		return
	}

	// 进入房间
	//flag, err = redis.EnterRoom(roomID, username)
	//if err != nil || flag != 1 {
	//	fmt.Println(err, flag)
	//	response.MakeFail(c, "enter Room failure")
	//	return
	//}
	// write to mysql
	room, err := model.CreateRoom(roomID, reqRoom.Name)
	if err != nil {
		response.MakeFail(c, "房间创建错误")
		return
	}

	response.MakeSuccessString(c, room.ID)
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
	response.MakeSuccessString(c, resRoom.Name)
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

	// 判断此用户是否已在房间中
	flag, _ := redis.UserExistRoom(username)
	if flag == 1 { // 在房间中,离开
		//获取所在房间的ID
		oldRoomID, _ := redis.GetUserInRoom(username)
		_, err := redis.LeaveRoom(oldRoomID, username)
		if err != nil {
			response.MakeFail(c, "leave Room failure")
		}
	}

	// 判断房间是否存在
	flag, err := redis.RoomExists(roomID)
	if err != nil || flag != 1 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}

	// 进入房间
	flag, err = redis.EnterRoom(roomID, username)
	if err != nil || flag != 1 {
		response.MakeFail(c, "enter Room failure")
		return
	}
	response.MakeSuccessString(c, "enter the Room success")
}

// LeaveRoom user leaven room
func LeaveRoom(c *gin.Context) {
	username := c.MustGet("username").(string)

	// 判断此用户是否已在房间中
	flag, err := redis.UserExistRoom(username)
	if err != nil || flag != 1 {
		response.MakeFail(c, "leave Room failure")
		return
	}

	//获取所在房间的ID
	oldRoomID, _ := redis.GetUserInRoom(username)
	flag, err = redis.LeaveRoom(oldRoomID, username)
	if err != nil || flag != 1 {
		response.MakeFail(c, "leave Room failure")
		return
	}
	response.MakeSuccessString(c, "left the room")
	return
}

// RoomAllUser get room all user list
func RoomAllUser(c *gin.Context) {
	roomID := c.Param("roomid")
	if len(roomID) <= 0 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}
	// 房间是否存在
	flag, err := redis.RoomExists(roomID)
	if err != nil || flag != 1 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}

	roomUser, err := redis.GetRoomAllUser(roomID)
	if err != nil {
		response.MakeFail(c, "get users failure")
	}

	response.MakeSuccessJSON(c, roomUser)
}
