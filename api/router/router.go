package router

import (
	"chat-room-go/api/router/handlers"
	"chat-room-go/api/router/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// Load load the middlewares, routers
func Load(e *fiber.App) *fiber.App {

	e.Use(compress.New())

	// user router
	user := e.Group("/")
	{
		user.Post("/user", handlers.CreateUser)                    // Create user
		user.Get("/userLogin", handlers.UserLogin)                 // Logs user into the system
		user.Get("/user/:username", cache.New(), handlers.GetUser) // Get user by user name
	}
	// room router
	room := e.Group("/")
	{
		room.Post("/room", middleware.JWTAuth(), handlers.CreateRoom)             // Create a new room
		room.Put("/room/:roomid/enter", middleware.JWTAuth(), handlers.EnterRoom) // Enter a room
		room.Put("/roomLeave", middleware.JWTAuth(), handlers.LeaveRoom)          // Leave a root
		room.Get("/room/:roomid", cache.New(), handlers.GetOneRoomInfo)           // Get the room info
		room.Get("/room/:roomid/users", handlers.RoomAllUser)                     // Get user list in a room, only username in list
		room.Post("/roomList", handlers.GetRoomList)                              // Get the room list
	}

	// message router
	message := e.Group("/")
	{
		message.Post("/message/send", middleware.JWTAuth(), handlers.SendMessage)        // After enter a room, the user can send the message to the current room.
		message.Post("/message/retrieve", middleware.JWTAuth(), handlers.GetMessageList) // After enter a room, the user can retrieve the message in the current room
	}

	// cluster router
	cluster := e.Group("/")
	{
		cluster.Post("/updateCluster", handlers.UpdateCluster)
		cluster.Get("/checkCluster", handlers.CheckCluster)
	}

	return e
}
