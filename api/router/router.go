package router

import (
	"chat-room-go/api/router/handlers"
	"chat-room-go/api/router/middleware"
	"github.com/gin-gonic/gin"
)

// Load load the middlewares, routers
func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	//e.Use(gin.Recovery())
	//e.Use(gin.Logger())
	e.Use(mw...)

	// user router
	user := e.Group("/")
	{
		user.POST("/user", handlers.CreateUser)                             // Create user
		user.GET("/userLogin", handlers.UserLogin)                          // Logs user into the system
		user.GET("/user/:username", middleware.JWTAuth(), handlers.GetUser) // Get user by user name
	}
	// room router
	room := e.Group("/", middleware.JWTAuth())
	{
		room.POST("/room", handlers.CreateRoom)               // Create a new room
		room.PUT("/room/:roomid/enter", handlers.EnterRoom)   // Enter a room
		room.PUT("/roomLeave", handlers.LeaveRoom)            // Leave a root
		room.GET("/room/:roomid", handlers.GetOneRoomInfo)    // Get the room info
		room.GET("/room/:roomid/users", handlers.RoomAllUser) // Get user list in a room, only username in list
		room.POST("/roomList", handlers.GetRoomList)          // Get the room list
	}

	// message router
	message := e.Group("/", middleware.JWTAuth())
	{
		message.POST("/message/send", handlers.SendMessage)        // After enter a room, the user can send the message to the current room.
		message.POST("/message/retrieve", handlers.GetMessageList) // After enter a room, the user can retrieve the message in the current room
	}

	// cluster router
	cluster := e.Group("/")
	{
		cluster.POST("/updateCluster", handlers.UpdateCluster)
		cluster.GET("/checkCluster", handlers.CheckCluster)
	}
	return e

}
