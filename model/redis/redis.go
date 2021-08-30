package redis

import (
	"github.com/golang/glog"
)

var (
	rs RedisInterface
)

// InitRedis 初始化Redis数据库
func InitRedis() (err error) {

	rs, err = NewRedis()
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func InitRedisSentinel(host []string, masterName string, password string) (err error) {
	rs, err = NewRedisSentinel(host, masterName, password)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func CloseRedis() {
	rs = nil
}

//func GetRedisMasterIP() (string, error) {
//	if rs == nil {
//		return "",errors.New("no init redis")
//	}
//
//}
