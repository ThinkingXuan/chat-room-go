package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
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
	Full(key string) bool
}

// RedisInterface  redis所有操作的接口
type RedisInterface interface {
	ListInterface
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

	// 要求RRedis结构体实现返回的接口中所有的方法！
	redisObj := &RRedis{
		maxIdle:        100,
		maxActive:      130,
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
