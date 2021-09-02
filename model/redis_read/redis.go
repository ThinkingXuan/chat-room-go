package redis_read

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

func CloseRedis() {
	rs = nil
}

//GetRedisMasterIP 获取Redis Sentinel的Master节点IP
func GetRedisMasterIP() string {
	return curMasterNodeIP
}
