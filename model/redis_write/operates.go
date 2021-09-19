package redis_write

import (
	"chat-room-go/util"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math"
	"time"
)

type RRedis struct {
	redisCli       *redis.Pool
	masterAddr     string
	maxIdle        int
	maxActive      int
	maxIdleTimeout time.Duration
	maxTimeout     time.Duration
	lazyLimit      bool
	maxSize        int
}

// 从池里获取连接 ———— 单独的方法
func (r *RRedis) getRedisConn() redis.Conn {
	rc := r.redisCli.Get()
	// // 用完后将连接放回连接池
	// defer rc.Close()
	return rc
}

// GetAllKeys 获取所有keys
func (r *RRedis) GetAllKeys() []string {

	rc := r.getRedisConn()
	defer rc.Close()

	keys, err := redis.Strings(rc.Do("KEYS", "*"))
	if err != nil {
		return make([]string, 0)
	}
	return keys
}

func (r *RRedis) Get(key string, timeout int) (string, error) {

	start := time.Now()

	for {
		res, err := r.GetNoWait(key)
		if err != nil {
			return "", err
		} else if res == "" {
			if timeout != -1 {
				lasted := time.Now().Sub(start)
				if r.maxTimeout > lasted {
					t1 := r.maxTimeout
					t2 := time.Duration(timeout)*time.Second - lasted
					time.Sleep(time.Duration(math.Min(float64(t1), float64(t2))))
				} else {
					return "", errors.New("GET timeout")
				}
			} else {
				time.Sleep(r.maxTimeout)
			}
		} else {
			return res, nil
		}
	}
}

func (r *RRedis) GetNoWait(key string) (string, error) {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.String(rc.Do("LPOP", key))

	if err != nil {
		return "", err
	}
	return res, nil
}

func (r *RRedis) Put(key string, value string, timeout int) (int, error) {

	start := time.Now()

	for {
		res, err := r.PutNoWait(key, value)

		if err != nil {
			return 0, err
		} else if res == -1 {
			if timeout != -1 {
				lasted := time.Now().Sub(start)
				if r.maxTimeout > lasted {
					t1 := r.maxTimeout
					t2 := time.Duration(timeout)*time.Second - lasted
					time.Sleep(time.Duration(math.Min(float64(t1), float64(t2))))
				} else {
					return 0, errors.New("PUT timeout")
				}
			} else {
				time.Sleep(r.maxTimeout)
			}

		} else {
			return res, nil
		}

	}
}

func (r *RRedis) PutNoWait(key string, value string) (int, error) {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("RPUSH", key, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) QSize(key string) int {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("LLEN", key))
	if err != nil {
		return -1
	}

	return res
}

func (r *RRedis) Empty(key string) bool {

	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("LLEN", key))
	if err != nil {
		return false
	}
	if res == 0 {
		return true
	}
	return false
}

func (r *RRedis) HPut(key string, field string, value interface{}) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("HSET", key, field, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) HGet(key string, field string) (interface{}, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.String(rc.Do("HGET", key, field))

	if err != nil {
		return "", err
	}
	return res, nil
}

func (r *RRedis) HDel(key string, field string) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("HDEL", key, field))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) HExists(key string, field string) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("HEXISTS", key, field))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) SPut(key string, value interface{}) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("SADD", key, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) SDel(key string, value string) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("SREM", key, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) SExists(key string, value string) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("SISMEMBER", key, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) SLen(key string) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("SCARD", key))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RRedis) SGetAll(key string) ([]string, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	// 数据量在百万以下，换成sscan并不能体现优势
	res, err := redis.Strings(rc.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RRedis) SGETScanAll(key string) ([]string, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	var res []string
	iter := 0
	for {
		// SSCAN key cursor [MATCH pattern] [COUNT count]
		// cursor - 游标。 每次查询会自动返回下次查询的游标
		// pattern - 匹配的模式。
		// count - 指定从数据集里返回多少元素，默认值为 10 。
		arr, err := redis.Values(rc.Do("SSCAN", key, iter, "count", 1000))
		if err != nil {
			return res, fmt.Errorf("error retrieving keys,%s", err)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		res = append(res, k...)

		if iter == 0 {
			break
		}
	}
	return res, nil
}

// ZsPUT ZSet插入
func (r *RRedis) ZsPUT(key string, score int64, value interface{}) (int, error) {
	rc := r.getRedisConn()
	defer rc.Close()

	res, err := redis.Int(rc.Do("ZADD", key, score, value))
	if err != nil {
		return 0, err
	}
	return res, nil
}

// ZsRange ZSet分页遍历
func (r *RRedis) ZsRange(key string, index, size int) (res []string, err error) {
	rc := r.getRedisConn()
	defer rc.Close()
	// start index
	start := util.IndexToPage(index, size)
	// stop index
	stop := start + size - 1

	if stop < 0 {
		return []string{}, nil
	}

	res, err = redis.Strings(rc.Do("ZRANGE", key, start, stop))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ZsRevRange ZSet分页倒序遍历
func (r *RRedis) ZsRevRange(key string, index, size int) (res []string, err error) {

	rc := r.getRedisConn()
	defer rc.Close()
	// start index
	start := util.IndexToPage(index, size)
	// stop index
	stop := start + size - 1

	if stop < 0 {
		return []string{}, nil
	}

	res, err = redis.Strings(rc.Do("ZREVRANGE", key, start, stop))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RRedis) CreateRoomAndRoomInfo(roomID, roomName string) error {
	rc := r.getRedisConn()
	defer rc.Close()
	_ = rc.Send("ZADD", RoomsKey, util.GetSnowflakeInt2(), roomID+"#"+roomName)
	_ = rc.Send("HSET", RoomInfoKey, roomID, roomName)
	rc.Flush()
	v, _ := rc.Receive()

	if  v.(int64) != 1 {
		return errors.New("create room and room info error")
	}
	return nil
}
