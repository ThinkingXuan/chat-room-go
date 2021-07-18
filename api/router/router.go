package router

import (
	"chat-room-go/api/router/handlers"
	"github.com/gin-gonic/gin"
)

// Load load the middlewares, routers
func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(mw...)

	// user router
	user := e.Group("/")
	{
		user.POST("/user", handlers.CreateUser)       // Create user
		user.GET("/userLogin", handlers.UserLogin)    // Logs user into the system
		user.GET("/user/:username", handlers.GetUser) // Get user by user name
	}
	// room router
	room := e.Group("/")
	{
		room.POST("/room")              // Create a new room
		room.PUT("/room/:roomid/enter") // Enter a room
		room.PUT("/roomLeave")          // Leave a root
		room.GET("/room/:roomid")       // get the room info
		room.GET("/room/:roomid/users") // Get user list in a room, only username in list
		room.POST("/roomList")          // get the room list
	}

	// message router
	message := e.Group("/")
	{
		message.POST("/message/send")     // After enter a room, the user can send the message to the current room.
		message.POST("/message/retrieve") //After enter a room, the user can retrieve the message in the current room
	}

	return e
}
