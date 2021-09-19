package benchmark

import (
	"chat-room-go/model/redis_write"
	"chat-room-go/util"
	"fmt"
	"testing"
)

func BenchmarkRedisHsetInterface(b *testing.B) {

}

func TestRedisConnect(t *testing.T) {
	redisCLi, err := redis_write.ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	for i := 0; i < 100000; i++ {
		err = redisCLi.CreateRoomAndRoomInfo(util.GetSnowflakeID2(),util.GetSnowflakeID2())
		if err != nil {
			t.Log(err)
		}
	}
}