package redis

import (
	"github.com/golang/glog"
)

var (
	rs RedisInterface
)

//// InitRedis 初始化Redis数据库
//func InitRedis() (err error) {
//	rs, err = NewRedis()
//	if err != nil {
//		glog.Error(err)
//		return err
//	}
//	return nil
//}

// InitRedisSentinel 初始化带Sentinel的Redis接口
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

//GetRedisMasterIP 获取Redis Sentinel的Master节点IP
func GetRedisMasterIP() string {
	return curMasterNodeIP
}
