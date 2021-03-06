package main

import (
	"chat-room-go/api/router"
	"chat-room-go/config"
	"chat-room-go/model"
	"chat-room-go/model/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"strconv"
)

var cfgPath = "config/conf/config.yaml"

func main() {

	err := config.InitConfig(cfgPath)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	// set gin run mode
	gin.SetMode(viper.GetString("runmode"))
	gin.ForceConsoleColor()
	//初始化
	if err := model.InitSQLite(); err != nil {
		glog.Error(err)
		panic("数据库初始化失败")
	}
	defer model.Close()

	maxOpenConn, _ := strconv.Atoi(viper.GetString("gorm.max_open_conn"))
	maxIdleConn, _ := strconv.Atoi(viper.GetString("gorm.max_idle_conn"))
	fmt.Println(maxOpenConn)
	fmt.Println(maxIdleConn)

	//模型绑定
	model.InitDBTable(maxOpenConn, maxIdleConn)

	// 初始化 Redis
	if err := redis.InitRedis(); err != nil {
		glog.Error(err)
		panic("Redis初始化失败")
	}

	//配置路由和中间件
	r := router.Load(gin.New())

	// 运行gin服务
	if err := r.Run(viper.GetString("url")); err != nil {
		glog.Info(err)
	}

}
