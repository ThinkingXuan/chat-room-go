package main

import (
	"chat-room-go/api/router"
	"chat-room-go/config"
	"chat-room-go/internal/run"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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

	// init leveldb
	//err = myleveldb.InitLevelDB()
	//if err != nil {
	//	glog.Error(err)
	//}

	// init gin engine
	runMode := viper.GetString("runmode")
	g := gin.New()

	switch runMode {
	case "release":
	case "debug":
		g.Use(gin.Recovery())
		g.Use(gin.Logger())
	}
	// set gin run mode
	gin.SetMode(runMode)

	// init gin router and middleware
	r := router.Load(g, gzip.Gzip(gzip.DefaultCompression))

	// run gin service
	if err := r.Run(viper.GetString("url")); err != nil {
		glog.Info(err)
	}
}
