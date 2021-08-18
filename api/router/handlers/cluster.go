package handlers

import (
	"chat-room-go/api/router/response"
	"chat-room-go/api/router/rr"
	"chat-room-go/model/redis"
	"chat-room-go/model/redis/tool"
	"chat-room-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"sort"
	"strings"
)

var (
	host []string
)

func UpdateCluster(c *gin.Context) {

	var reqClusterIP rr.ReqClusterIP
	if err := c.ShouldBindJSON(&reqClusterIP); err != nil {
		response.MakeFail(c, "param err")
		return
	}

	// sort
	sort.Strings(reqClusterIP)

	localIP := util.GetIp()
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

	// 赋予权限脚本
	_, err := util.ExecShell("sudo chmod +x config/script/rediscluster/redismasl.sh")
	_, err = util.ExecShell("sudo chmod +x config/script/rediscluster/sentinel.sh")

	// redis master/slave peer script start
	if peerNum == 0 {
		_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s", "master"))
	} else {
		_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s %s", "slave", reqClusterIP[0]))
	}

	// exec redis sentinel script
	_, err = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/sentinel.sh %s", reqClusterIP[0]))

	if err != nil {
		response.MakeFail(c, "redis or sentinel start failure")
		return
	}

	// host
	host = reqClusterIP
	hostStr := ""
	port := viper.GetString("redis-sentinel.port")
	for i := 0; i < len(host); i++ {
		host[i] += ":" + port
		hostStr += "\n" + host[i]
	}
	// host write to file
	util.WriteWithFile("./ip-address", hostStr)

	masterName := viper.GetString("redis-sentinel.master_name")
	password := viper.GetString("redis-sentinel.password")

	// init redis sentinel client
	err = redis.InitRedisSentinel(host, masterName, password)
	if err != nil {
		response.MakeFail(c, "redis client start failure")
		return
	}

	response.MakeSuccessString(c, "success")
}

func CheckCluster(c *gin.Context) {

	masterName := viper.GetString("redis-sentinel.master_name")
	password := viper.GetString("redis-sentinel.password")

	// read a file ip address
	if len(host) <= 0 {
		hostString := util.ReadWithFile("./ip-address")
		host = strings.Split(hostString, "\n")
	}
	// init redis sentinel client
	err := redis.InitRedisSentinel(host, masterName, password)
	if err != nil {
		response.MakeFail(c, "redis client start failure")
		return
	}

	redis, err := tool.NewRedisSentinel(host, masterName, password)
	if err != nil || redis == nil {
		response.MakeFail(c, "redis start failure")
		return
	}
	response.MakeSuccessString(c, "success")
}
