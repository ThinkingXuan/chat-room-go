package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"chat-room-go/util"
	"github.com/gofiber/fiber/v2"
)

// CreateRoom Create a room
func CreateRoom(c *fiber.Ctx) error {
	//username := c.MustGet("username").(string)
	var reqRoom rr.ReqRoom
	if err := c.BodyParser(&reqRoom); err != nil {
		return response.MakeFail(c, "param err")
	}
	if len(reqRoom.Name) <= 0 {
		return response.MakeFail(c, "Invalid input")
	}

	// create a room id
	roomID := util.GetSnowflakeID2()

	// write to redis： crate a room zset
	flag, err := redis_write.CreateRoom(roomID, reqRoom.Name)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "create room err")
	}
	//	write to redis： crate a room info
	flag, err = redis_write.CreateRoomInfo(roomID, reqRoom.Name)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "create room err")

	}

	return response.MakeSuccessString(c, roomID)

}

// GetOneRoomInfo Get a room information by roomID
func GetOneRoomInfo(c *fiber.Ctx) error {
	roomID := c.Params("roomid")
	if len(roomID) <= 0 {
		return response.MakeFail(c, "param err")
	}

	roomName, err := redis_read.GetRoomInfo(roomID)
	if len(roomName) <= 0 || err != nil {
		return response.MakeFail(c, "Invalid Room ID")
	}
	return response.MakeSuccessString(c, roomName)
}

// GetRoomList Get room page list
func GetRoomList(c *fiber.Ctx) error {
	var reqPage rr.ReqPage
	if err := c.BodyParser(&reqPage); err != nil {
		return response.MakeFail(c, "param err")
	}

	if reqPage.PageSize < 0 || reqPage.PageIndex < 0 {
		return response.MakeFail(c, "param err")
	}

	rooms, err := redis_read.SelectRoomListPage(reqPage.PageIndex, reqPage.PageSize)
	if err != nil {
		return response.MakeFail(c, "select err")
	}
	return response.MakeSuccessJSON(c, rooms)
}

// EnterRoom user enter room
func EnterRoom(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	roomID := c.Params("roomid")
	// 参数校验
	if len(roomID) <= 0 {
		return response.MakeFail(c, "Invalid Room ID")
	}

	// 判断房间是否存在
	flag, err := redis_read.RoomExists(roomID)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "Invalid Room ID")
	}

	//获取所在房间的ID
	oldRoomID, _ := redis_read.GetUserInRoom(username)

	// 现在所在房间为要进入的房间
	if oldRoomID == roomID {
		return response.MakeSuccessString(c, "enter the Room success")
	}
	// 现在所在房间为不是自己要进入的房间
	if len(oldRoomID) > 0 {
		// 离开房间
		_, err := redis_write.LeaveRoom(oldRoomID, username)
		if err != nil {
			return response.MakeFail(c, "leave Room failure")
		}
	}

	// 进入此房间
	flag, err = redis_write.EnterRoom(roomID, username)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "enter Room failure")
	}
	return response.MakeSuccessString(c, "enter the Room success")
}

// LeaveRoom user leaven room
func LeaveRoom(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	//获取所在房间的ID
	oldRoomID, _ := redis_read.GetUserInRoom(username)
	// 用户不在房间
	if len(oldRoomID) <= 0 {
		return response.MakeFail(c, "leave Room failure")
	}
	flag, err := redis_write.LeaveRoom(oldRoomID, username)
	if err != nil || flag != 1 {
		return response.MakeFail(c, "leave Room failure")
	}
	return response.MakeSuccessString(c, "left the room")
}

// RoomAllUser get room all user list
func RoomAllUser(c *fiber.Ctx) error {
	roomID := c.Params("roomid")
	if len(roomID) <= 0 {
		return response.MakeFail(c, "Invalid Room ID")
	}

	roomUser, err := redis_read.GetRoomAllUser(roomID)
	if err != nil || len(roomUser) < 0 {
		return response.MakeFail(c, "get users failure")
	}

	return response.MakeSuccessJSON(c, roomUser)
}
