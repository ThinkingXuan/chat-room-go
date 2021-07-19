package redis

import (
	"fmt"
	"testing"
)

func TestRedisConnect(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	fmt.Println(redisCLi)
}

func TestRedisWrite(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	// 获取所有keys
	allKeysLst := redisCLi.GetAllKeys()
	fmt.Print("key>>> ", allKeysLst)
}
