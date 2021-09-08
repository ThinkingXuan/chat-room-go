package redis

import (
	"chat-room-go/util"
	"fmt"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
	"time"
)

var curMasterNodeIP string

// RedisFatherInterface 父级interface
type RedisFatherInterface interface {
	// 获取所有keys
	GetAllKeys() []string
}

// ListInterface 操作list接口
type ListInterface interface {
	// "继承"父类的所有方法
	RedisFatherInterface
	GetNoWait(key string) (string, error) // ~di: interface{}类型、数据压缩
	Get(key string, timeout int) (string, error)
	Put(key string, value string, timeout int) (int, error)
	PutNoWait(key string, value string) (int, error)
	QSize(key string) int
	Empty(key string) bool
	//Full(key string) bool
}

// HashInterface 操作hash的接口
type HashInterface interface {
	HPut(key string, field string, value interface{}) (int, error)
	HGet(key string, field string) (interface{}, error)
	HDel(key string, field string) (int, error)
	HExists(key string, field string) (int, error)
}

// SETInterface 操作SET的接口
type SETInterface interface {
	SPut(key string, value interface{}) (int, error)
	SDel(key string, value string) (int, error)
	SExists(key string, value string) (int, error)
	SLen(key string) (int, error)
	SGetAll(key string) ([]string, error)
	SGETScanAll(key string) ([]string, error)
}

// ZSETInterface 操作zset的接口
type ZSETInterface interface {
	ZsPUT(key string, score int64, value interface{}) (int, error)
	ZsRange(key string, index, size int) ([]string, error)
	ZsRevRange(key string, index, size int) ([]string, error)
}

// RedisInterface  redis所有操作的接口
type RedisInterface interface {
	ListInterface
	HashInterface
	SETInterface
	ZSETInterface
}

// NewRedisSentinel create a redis sentinel
func NewRedisSentinel(hosts []string, masterName string, password string) (RedisInterface, error) {
	return ProduceRedisSentinel(hosts, masterName, password, 0, 100, true)
}

// ProduceRedisSentinel 工厂函数，要求对应的结构体必须实现 RedisInterface 中的所有方法
// 如果只想实现某一些方法，就返回"有这些方法的结构体"就好了
func ProduceRedisSentinel(hosts []string, masterName string, password string, db, maxSize int, lazyLimit bool) (RedisInterface, error) {

	maxActive, _ := strconv.Atoi(viper.GetString("redis.max_active_conn"))
	maxIdle, _ := strconv.Atoi(viper.GetString("redis.max_idle_conn"))

	sntnl := &sentinel.Sentinel{
		Addrs:      hosts,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	// 获取sentinel中的主节点的地址
	masterAddr, err := sntnl.MasterAddr()
	if err != nil {
		return nil, err
	}
	curMasterNodeIP = masterAddr
	go CheckMasterStatus(sntnl)

	// 要求RRedis结构体实现返回的接口中所有的方法！
	redisObj := &RRedis{
		masterAddr:     masterAddr,
		maxIdle:        maxIdle,
		maxActive:      maxActive,
		maxIdleTimeout: time.Duration(60) * time.Second,
		maxTimeout:     time.Duration(30) * time.Second,
		lazyLimit:      lazyLimit,
		maxSize:        maxSize,
	}
	//

	// 建立连接池
	redisObj.redisCli = &redis.Pool{
		MaxIdle:     redisObj.maxIdle,
		MaxActive:   redisObj.maxActive,
		IdleTimeout: redisObj.maxIdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(
				"tcp",
				redisObj.masterAddr,
				redis.DialPassword(password),
				redis.DialDatabase(int(db)),
				redis.DialConnectTimeout(redisObj.maxTimeout),
				redis.DialReadTimeout(redisObj.maxTimeout),
				redis.DialWriteTimeout(redisObj.maxTimeout),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: CheckRedisRole,
	}

	return redisObj, nil
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}
}

func CheckMasterStatus(sentinel *sentinel.Sentinel) {
	masterName := viper.GetString("redis-sentinel.master_name")
	password := viper.GetString("redis-sentinel.password")
	// read a file ip address
	hostString := util.ReadWithFile("./ip-address")
	host := strings.Split(hostString, "\n")

	for {
		time.Sleep(time.Second)
		//log.Println("curIP", curMasterNodeIP)
		masterNodeIP, _ := sentinel.MasterAddr()
		//log.Println("masterIP", masterNodeIP)
		if curMasterNodeIP != masterNodeIP {
			log.Println("cluster err!!!")
			// close redis
			CloseRedis()
			log.Println("restart cluster !!!")
			// init redis sentinel client
			err := InitRedisSentinel(host, masterName, password)
			if err != nil {
				continue
			}
			// 退出
			return
		}
		//log.Println("cluster healthy!!")
	}
}
