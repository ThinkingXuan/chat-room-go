package router

import (
	"chat-room-go/api/router/handlers"
	"chat-room-go/api/router/middleware"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

// Load load the middlewares, routers
func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	e.Use(mw...)

	// gin cache
	store := persistence.NewInMemoryStore(time.Second * 10)

	// user router
	user := e.Group("/")
	{
		user.POST("/user", handlers.CreateUser)    // Create user
		user.GET("/userLogin", handlers.UserLogin) // Logs user into the system
		//cache page
		user.GET("/user/:username", cache.CachePage(store, time.Second*60, handlers.GetUser)) // Get user by user name
	}
	// room router
	room := e.Group("/")
	{
		room.POST("/room", middleware.JWTAuth(), handlers.CreateRoom)             // Create a new room
		room.PUT("/room/:roomid/enter", middleware.JWTAuth(), handlers.EnterRoom) // Enter a room
		room.PUT("/roomLeave", middleware.JWTAuth(), handlers.LeaveRoom)          // Leave a root
		//cache page
		room.GET("/room/:roomid", cache.CachePage(store, time.Second*60, handlers.GetOneRoomInfo)) // Get the room info
		room.GET("/room/:roomid/users", handlers.RoomAllUser)                                      // Get user list in a room, only username in list
		room.POST("/roomList", handlers.GetRoomList)                                               // Get the room list
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
