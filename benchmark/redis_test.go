package benchmark

import (
	"chat-room-go/model/redis_write"
	"chat-room-go/util"
	"testing"
)

func BenchmarkRedisHsetInterface(b *testing.B) {

}

func TestRedisPipeline(t *testing.T) {
	c, _ := redis_write.NewRedis()
	err := c.CreateRoomAndRoomInfo(util.GetSnowflakeID2(), util.GetSnowflakeID2())
	t.Log(err)
}
