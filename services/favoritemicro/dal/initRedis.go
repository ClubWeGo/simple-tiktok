package dal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Dialer: dal.RedisDial,
		// SSH不支持超时设置，在这里禁用
		//ReadTimeout:  -1,
		//WriteTimeout: -1,
		Password: "",
		DB:       0,
	})
	pong, err := Redis.Ping(context.Background()).Result()
	fmt.Println(pong, err)
}
