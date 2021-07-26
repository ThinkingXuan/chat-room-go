package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis"
	"chat-room-go/util"
	"github.com/gin-gonic/gin"
)

// CreateRoom Create a room
func CreateRoom(c *gin.Context) {
	//username := c.MustGet("username").(string)
	var reqRoom rr.ReqRoom
	if err := c.ShouldBindJSON(&reqRoom); err != nil {
		response.MakeFail(c, "param err")
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
	roomID := util.GetSnowflakeID2()

	// write to redis： crate a room zset
	flag, err := redis.CreateRoom(roomID, reqRoom.Name)
	if err != nil || flag != 1 {
		response.MakeFail(c, "create room err")
		return
	}
	//	write to redis： crate a room info
	flag, err = redis.CreateRoomInfo(roomID, reqRoom.Name)
	if err != nil || flag != 1 {
		response.MakeFail(c, "create room err")
		return
	}

	//_, err = model.CreateRoom(roomID, reqRoom.Name)
	//if err != nil {
	//	response.MakeFail(c, "create room err")
	//	return
	//}

	// response
	response.MakeSuccessString(c, roomID)

	//异步写入MySQL
	//pool.Work(roomID, reqRoom.Name)
	//go model.CreateAsyncRoom(roomID, reqRoom.Name)
}

// GetOneRoomInfo Get a room information by roomID
func GetOneRoomInfo(c *gin.Context) {
	roomID := c.Param("roomid")
	if len(roomID) <= 0 {
		response.MakeFail(c, "param err")
		return
	}

	roomName, err := redis.GetRoomInfo(roomID)
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

	rooms, err := redis.SelectRoomListPage(reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		response.MakeFail(c, "select err")
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

	//// 判断房间是否存在
	//flag, err := redis.RoomExists(roomID)
	//if err != nil || flag != 1 {
	//	response.MakeFail(c, "Invalid Room ID")
	//	return
	//}

	////获取用户现在所在房间
	//oldRoomID, _ := redis.GetUserInRoom(username)
	//
	//// 现在所在房间为要进入的房间
	//if oldRoomID == roomID {
	//	response.MakeFail(c, "user already in room")
	//	return
	//}
	//
	//// 现在所在房间不是要进入的房间
	//if len(oldRoomID) > 0 {
	//	// 离开房间
	//	_, err := redis.LeaveRoom(oldRoomID, username)
	//	if err != nil {
	//		response.MakeFail(c, "leave Room failure")
	//	}
	//}

	// 判断房间是否存在
	flag, err := redis.RoomExists(roomID)
	if err != nil || flag != 1 {
		response.MakeFail(c, "Invalid Room ID")
		return
	}

	//获取所在房间的ID
	oldRoomID, _ := redis.GetUserInRoom(username)

	// 现在所在房间为要进入的房间
	if oldRoomID == roomID {
		response.MakeSuccessString(c, "enter the Room success")
		return
	}
	// 现在所在房间为不是自己要进入的房间
	if len(oldRoomID) > 0 {
		// 离开房间
		_, err := redis.LeaveRoom(oldRoomID, username)
		if err != nil {
			response.MakeFail(c, "leave Room failure")
		}
	}

	// 进入此房间
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
	//flag, err := redis.UserExistRoom(username)
	//if err != nil || flag != 1 {
	//	response.MakeFail(c, "leave Room failure")
	//	return
	//}

	//获取所在房间的ID
	oldRoomID, _ := redis.GetUserInRoom(username)
	// 用户不在房间
	if len(oldRoomID) < 0 {
		response.MakeFail(c, "leave Room failure")
		return
	}
	flag, err := redis.LeaveRoom(oldRoomID, username)
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
	//flag, err := redis.RoomExists(roomID)
	//if err != nil || flag != 1 {
	//	response.MakeFail(c, "Invalid Room ID")
	//	return
	//}

	roomUser, err := redis.GetRoomAllUser(roomID)
	if err != nil || len(roomUser) < 0 {
		response.MakeFail(c, "get users failure")
	}

	response.MakeSuccessJSON(c, roomUser)
}
