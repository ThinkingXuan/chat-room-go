package main

import (
	"chat-room-go/api/router"
	"chat-room-go/config"
	"chat-room-go/internal/run"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var cfgPath = "config/conf/config.yaml"

func main() {

	// init config
	err := config.InitConfig(cfgPath)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	// start redis_write sentinel and client connection
	run.StartRedisSentinelAndClientConnection()

	// init fiber router and middleware
	r := router.Load(fiber.New())

	// run fiber service
	if err := r.Listen(viper.GetString("url")); err != nil {
		glog.Info(err)
	}
}
