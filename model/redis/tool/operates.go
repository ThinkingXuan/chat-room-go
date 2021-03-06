package tool

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"math"
	"time"
)

type RRedis struct {
	redisCli       *redis.Pool
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

	res, err := redis.Strings(rc.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}
