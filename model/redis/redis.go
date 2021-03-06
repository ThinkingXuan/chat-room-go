package redis

import (
	"chat-room-go/model/redis/tool"
	"github.com/golang/glog"
)

var (
	rs tool.RedisInterface
)

// InitRedis 初始化Redis数据库
func InitRedis() (err error) {

	rs, err = tool.NewRedis()
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}
