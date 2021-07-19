package tool

import (
	"fmt"
	"github.com/magiconair/properties/assert"
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

	v, err := redisCLi.HPut("rooms", "12312324", "fly2")
	if err != nil {
		t.Log(err)
	}
	t.Log(v)

	i, err := redisCLi.HGet("rooms", "12312324")
	if err != nil {
		t.Log(err)
	}
	t.Log(i.(string))

	v2, err := redisCLi.HDel("rooms", "12312324")
	if err != nil {
		t.Log(err)
	}
	t.Log(v2)
}

func TestHSET(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.HPut("rooms", "123112323432123422324", "fly2")
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, v, 1)

	v2, err := redisCLi.HPut("rooms", "123112323432123422324", "fly3")
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, v2, 0)
}

func TestHGET(t *testing.T) {
	TestHSET(t)
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}

	i, err := redisCLi.HGet("rooms", "1231123123422324")
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, i.(string), "fly2")
}

func TestHDEL(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v2, err := redisCLi.HDel("rooms", "12312324")
	if err != nil {
		t.Log(err)
	}
	t.Log(v2)
}

func TestHExists(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v2, err := redisCLi.HExists("rooms", "123123234244")
	if err != nil {
		t.Log(err)
	}
	t.Log(v2)
}

func TestRRedis_SPut(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v1, err := redisCLi.SPut("123123234244", "youxuan34543522")
	if err != nil {
		t.Log(err)
	}
	t.Log(v1)
}

func TestRRedis_SGetAll(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.SGetAll("123123234244")
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}

func TestRRedis_SExists(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.SExists("123123234244", "youxuan34543522")
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}

func TestRRedis_SLen(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.SLen("123123234244")
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}

func TestRRedis_SDel(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.SDel("123123234244", "youxuan")
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}
