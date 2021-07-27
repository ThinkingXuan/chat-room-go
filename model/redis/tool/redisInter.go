package tool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

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

// NewRedis create a redis connect
func NewRedis() (RedisInterface, error) {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	return ProduceRedis(host, port, password, 0, 100, true)
}

// ProduceRedis 工厂函数，要求对应的结构体必须实现 RedisInterface 中的所有方法
// 如果只想实现某一些方法，就返回"有这些方法的结构体"就好了
func ProduceRedis(host, port, password string, db, maxSize int, lazyLimit bool) (RedisInterface, error) {

	maxActive, _ := strconv.Atoi(viper.GetString("redis.max_active_conn"))
	maxIdle, _ := strconv.Atoi(viper.GetString("redis.max_idle_conn"))

	// 要求RRedis结构体实现返回的接口中所有的方法！
	redisObj := &RRedis{
		maxIdle:        maxIdle,
		maxActive:      maxActive,
		maxIdleTimeout: time.Duration(60) * time.Second,
		maxTimeout:     time.Duration(30) * time.Second,
		lazyLimit:      lazyLimit,
		maxSize:        maxSize,
	}
	// 建立连接池
	redisObj.redisCli = &redis.Pool{
		MaxIdle:     redisObj.maxIdle,
		MaxActive:   redisObj.maxActive,
		IdleTimeout: redisObj.maxIdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(
				"tcp",
				host+":"+port, // address
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
	}

	return redisObj, nil
}
