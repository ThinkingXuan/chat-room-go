package run

import (
	"chat-room-go/model/redis_read"
	"chat-room-go/model/redis_write"
	"chat-room-go/util"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type RedisConnect struct {
	host       []string
	masterName string
	password   string
}

// StartRedisSentinelAndClientConnection 启动Redis哨兵集群，并使用客户端连接。主要用于宕机重启时。
// 判断是否时宕机重启的方式时判断程序根目录是否存在ip-address文件
func StartRedisSentinelAndClientConnection() {

	// 第一次启动
	if !isPreviousRun() {
		return
	}
	// 已启动过
	// 读取配置文件
	rc, _ := ReadRedisSentinelConfig()
	// 启动redis哨兵和集群
	StartRedisSentinel(rc)
	// 初始化客户端连接
	startClientConnectionAt(rc)
}

// 判断是否之前启动过，true: 已经启动过  false: 第一次启动
func isPreviousRun() bool {
	ipAddressFile, _ := os.Stat("./ip-address") // 之前已经启动过
	if ipAddressFile != nil {
		return true
	}
	return false
}

// StartRedisSentinel 启动Redis哨兵集群
func StartRedisSentinel(rc RedisConnect) {
	// 连接还存活的redis集群，目的是启动集群，通过获取master节点ip地址
	_ = redis_write.InitRedisSentinel(rc.host, rc.masterName, rc.password)
	// 延迟2s
	time.Sleep(time.Second * 2)
	// 获取master节点地址  ip:端口
	masterIP := redis_write.GetRedisMasterIP()
	// 去除端口号
	masterIP = strings.Split(masterIP, ":")[0]

	// 机器完全宕机后，重新生成集群
	if masterIP == "" {
		hosts := getRedisSentinelHosts()
		masterIP = strings.Split(hosts[0], ":")[0]
		// todo
	}

	// 启动集群和哨兵(因为，redis和redis-sentinel没有设置开机启动，所以需要用命令启动)
	RunSlaveRedisClusterAndSentinel(masterIP)

}

// startClientConnection 客户端连接
func startClientConnectionAt(rc RedisConnect) {
	// 延迟2s
	time.Sleep(time.Second * 2)
	// 初始化redis的读和写的客户端连接
	_ = StartRedisReadConnection()
	_ = StartRedisWriteConnection(rc)
}

func StartRedisWriteConnection(rc RedisConnect) error {
	return redis_write.InitRedisSentinel(rc.host, rc.masterName, rc.password)

}

func StartRedisReadConnection() error {
	return redis_read.InitRedis()
}

// ReadRedisSentinelConfig 读取Redis-sentinel有关的配置
func ReadRedisSentinelConfig() (RedisConnect, error) {
	var rc RedisConnect

	hosts := getRedisSentinelHosts()
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

func getRedisSentinelHosts() []string {
	hostString := util.ReadWithFile("./ip-address")
	hosts := strings.Split(hostString, "\n")
	return hosts
}

// RunSlaveRedisClusterAndSentinel 启动集群
func RunSlaveRedisClusterAndSentinel(masterIP string) {
	_, _ = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/redismasl.sh %s %s", "slave", masterIP))
	_, _ = util.ExecShell(fmt.Sprintf("sudo sh config/script/rediscluster/sentinel.sh %s", masterIP))
}
