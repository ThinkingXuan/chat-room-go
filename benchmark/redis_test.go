package benchmark

import (
	"chat-room-go/model/redis/tool"
	"chat-room-go/util"
	"fmt"
	"testing"
)

func BenchmarkRedisHsetInterface(b *testing.B) {

	redisCLi, err := tool.ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}

	for n := 0; n < b.N; n++ {
		key := util.GetSnowflakeID2()
		redisCLi.HPut("test", key, "youxuan")
	}
}
