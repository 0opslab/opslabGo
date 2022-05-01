package main

import (
	"time"

	"github.com/go-redis/redis"
	logger "github.com/sirupsen/logrus"
)

var rediscli *redis.Client

func init_redis(address string) {
	rediscli = redis.NewClient(&redis.Options{
		Addr:     address,
		PoolSize: 100,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := rediscli.Ping().Result()
	logger.Info("init redis :", pong, err)
}

func getString(redisKey string) string {
	res, err := rediscli.Get(redisKey).Result()
	if err != nil {
		logger.Error("redis get Error:", err)
	}
	return res
	// res, _ := rediscli.Get(redisKey).Bytes()
	// if len(res) > 1 {
	// 	ress, err := FlateDecode(res)
	// 	if err != nil {
	// 		logger.Error("redis FlateDecode Error:", err)
	// 	}
	// 	return string(ress)
	// }
	// return ""
}

func setRedis(redisKey string, data interface{}, tt time.Duration) {
	tmp := toJosn(data)
	if len(tmp) > 0 {
		status := rediscli.Set(redisKey, string(tmp), tt)
		logger.Trace("set redis > key:", redisKey, "expire_time:", tt, "status:", status.Val())
	}
	// tmp := toJosn(data)
	// if len(tmp) > 0 {
	// 	s1, err := FlateEncode(tmp)
	// 	if err != nil {
	// 		logger.Error("set redis > key:", redisKey, "expire_time:", tt, "FlateEncodeByte:", err)
	// 	} else {
	// 		status := rediscli.Set(redisKey, s1, tt)
	// 		logger.Trace("set redis > key:", redisKey, "expire_time:", tt, "status:", status.Val())
	// 	}
	// }
}
