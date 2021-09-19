package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/internal/run"
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"chat-room-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"sort"
	"time"
)

func UpdateCluster(c *gin.Context) {

	var reqClusterIP rr.ReqClusterIP
	if err := c.ShouldBindJSON(&reqClusterIP); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	// sort
	sort.Strings(reqClusterIP)

	localIP := util.GetLocalIP()
	peerNum := -1
	for i := 0; i < len(reqClusterIP); i++ {
		if reqClusterIP[i] == localIP {
			peerNum = i
			break
		}
	}

	if len(localIP) <= 0 || peerNum == -1 {
		response.MakeFail(c, "ip address get failure")
		return
	}

	// 非master节点延迟2s
	if peerNum != 0 {
		time.Sleep(time.Second * 2)
	}

	// 赋予权限脚本
	_, err := util.ExecShell("sudo chmod +x config/script/rediscluster/redismasl.sh")
	_, err = util.ExecShell("sudo chmod +x config/script/rediscluster/sentinel.sh")

	// redis_write master/slave peer script start
	if peerNum == 0 {
		_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s", "master"))
	} else {
		_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s %s", "slave", reqClusterIP[0]))
	}

	// after delay 2s, start sentinel
	time.Sleep(time.Second * 2)

	// exec redis_write sentinel script
	_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/sentinel.sh %s", reqClusterIP[0]))

	if err != nil {
		response.MakeFail(c, "redis_write or sentinel start failure")
		return
	}

	// host
	host := reqClusterIP
	hostStr := ""
	port := viper.GetString("redis-sentinel.port")
	for i := 0; i < len(host); i++ {
		host[i] += ":" + port
		hostStr += "\n" + host[i]
	}
	// host write to file
	util.WriteWithFile("./ip-address", hostStr)

	response.MakeSuccessString(c, "success")
}

func CheckCluster(c *gin.Context) {

	rc, _ := run.ReadRedisSentinelConfig()

	// close redis
	redis_write.CloseRedis()
	redis_read.CloseRedis()

	// init redis_write write sentinel client
	err := run.StartRedisWriteConnection(rc)
	if err != nil {
		response.MakeFail(c, "redis_write client start failure")
		return
	}

	time.Sleep(1 * time.Second)

	// init redis_write read sentinel client
	err = run.StartRedisReadConnection()
	if err != nil {
		response.MakeFail(c, "redis_write client start failure")
		return
	}
	response.MakeSuccessString(c, "success")
}
