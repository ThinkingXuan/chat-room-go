package tool

import (
	"chat-room-go/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/magiconair/properties/assert"
	"sync"
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
	for i := 0; i < 100000; i++ {

		str := "youxuan" + util.GetSnowflakeID2()
		redisCLi.SPut("users", str)
		//t.Log(v1)
	}
	//v1, err := redisCLi.SPut("room", "youxuan34543522")
	//if err != nil {
	//	t.Log(err)
	//}
	//t.Log(v1)
}

func TestRRedis_SPut_Pipe(t *testing.T) {
	//redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	//if err != nil {
	//	fmt.Println("redis连接错误！err>>>", err.Error())
	//	return
	//}
	//redisCLi.SPutPipe()

	dialOption := redis.DialPassword("123456")
	conn, err := redis.Dial("tcp", "localhost:6379", dialOption)
	if err != nil {
		fmt.Println("conn redis failed, err:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 1000000; i++ {

		str := "youxuan" + util.GetSnowflakeID2()
		conn.Do("SADD", "name", str)
		//t.Log(v1)
	}
}

func TestRRedis_SGetAll(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 0; i < 50; i++ {
		go getall(&wg, t, redisCLi)
	}
	wg.Wait()
}

func TestRRedis_SGetScanAll(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go getScanall(&wg, t, redisCLi)
	}
	wg.Wait()

}

func getall(wg *sync.WaitGroup, t *testing.T, res RedisInterface) {
	_, err := res.SGetAll("users")
	if err != nil {
		fmt.Println(err)
	}
	//t.Log(v)
	wg.Done()
}

func getScanall(wg *sync.WaitGroup, t *testing.T, res RedisInterface) {
	_, err := res.SGETScanAll("users")
	if err != nil {
		fmt.Println(err)
	}
	//t.Log(v)
	wg.Done()
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
	v, err := redisCLi.SLen("users")
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

func TestSScan(t *testing.T) {
	dialOption := redis.DialPassword("123456")
	c, err := redis.Dial("tcp", "localhost:6379", dialOption)
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	defer c.Close()

	//users := []string{}

	s, _ := c.Do("SSCAN", "users", "1", "count", "3")
	fmt.Printf("%T\n", s)
	fmt.Printf("%T v:= %d\n", s.([]interface{})[0], s.([]interface{})[0].([]uint8)[0]-'0')
	fmt.Printf("%T\n", s.([]interface{})[1])

	//t, _ := redis.Ints(s.([]interface{})[0],nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func TestZrange(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.ZsRange("zyou", 2, 10)
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}

func TestZRevrange(t *testing.T) {
	redisCLi, err := ProduceRedis("127.0.0.1", "6379", "123456", 0, 100, true)
	if err != nil {
		fmt.Println("redis连接错误！err>>>", err.Error())
		return
	}
	v, err := redisCLi.ZsRevRange("zyou", 0, 10)
	if err != nil {
		t.Log(err)
	}
	t.Log(v)
}
