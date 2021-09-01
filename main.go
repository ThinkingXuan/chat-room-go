package main

import (
	"chat-room-go/api/router"
	"chat-room-go/config"
	"chat-room-go/model/redis"
	"chat-room-go/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strings"
)

var cfgPath = "config/conf/config.yaml"

type RedisConnect struct {
	host       []string
	masterName string
	password   string
}

func main() {

	err := config.InitConfig(cfgPath)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	// set gin run mode
	gin.SetMode(viper.GetString("runmode"))

	// init redis sentinel client
	ipAddressFile, _ := os.Stat("./ip-address") // 之前已经启动过
	if ipAddressFile != nil {
		// 连接还存活的redis集群，目的是启动集群，通过获取master节点ip地址
		rc, _ := ReadRedisConfig()
		err := redis.InitRedisSentinel(rc.host, rc.masterName, rc.password)
		if err != nil {
			panic(err)
		}
		// 获取master节点地址  ip:端口
		masterIP := redis.GetRedisMasterIP()
		// 去除端口号
		masterIP = strings.Split(masterIP, ":")[0]
		// 启动集群和哨兵
		startRedisClusterAndSentinel(masterIP)
	}

	// init router and middleware
	r := router.Load(gin.New())

	// run gin service
	if err := r.Run(viper.GetString("url")); err != nil {
		glog.Info(err)
	}

	//if runtime.GOOS == "windows" {
	//	if err := r.Run(viper.GetString("url")); err != nil {
	//		glog.Info(err)
	//	}
	//}
	//else if runtime.GOOS == "linux" {
	//	// 优雅停止服务器
	//	server := endless.NewServer(viper.GetString("url"), r)
	//	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGINT] = append(
	//		server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGINT],
	//		sendQuickMessage)
	//	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP] = append(
	//		server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP],
	//		sendQuickMessage)
	//	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM] = append(
	//		server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM],
	//		sendQuickMessage)
	//	if err := server.ListenAndServe(); err != nil {
	//		log.Println(err)
	//	}
	//}
}

func sendQuickMessage() {
	// 服务终止
	masterHost := util.ReadWithFile("./master-address")
	localIP := util.GetIp()
	masterIP := strings.Split(masterHost, ":")[0]

	// 死掉的机器不是自己
	if masterIP != localIP {
		return
	}
	// 死掉的机器是自己，给其他机器发送信号
	otherHostsString := util.ReadWithFile("./ip-address")
	hosts := strings.Split(otherHostsString, "\n")
	for i := 0; i < len(hosts); i++ {
		ipAddr := strings.Split(hosts[i], ":")[0]
		if ipAddr != localIP {
			_, err := http.Get(fmt.Sprintf("http://%s:8080/startCluster", ipAddr))
			if err != nil {
				log.Println(err)
			}
		}
	}

}

// ReadRedisConfig 读取Redis配置
func ReadRedisConfig() (RedisConnect, error) {
	var rc RedisConnect
	hostString := util.ReadWithFile("./ip-address")
	hosts := strings.Split(hostString, "\n")
	if len(hosts) == 3 {
		masterName := viper.GetString("redis-sentinel.master_name")
		password := viper.GetString("redis-sentinel.password")
		rc.host = hosts
		rc.masterName = masterName
		rc.password = password
	} else {
		return rc, errors.New("host err")
	}
	return rc, nil
}

func startRedisClusterAndSentinel(masterIP string) {
	util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s %s", "slave", masterIP))
	util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/sentinel.sh %s", masterIP))
}
