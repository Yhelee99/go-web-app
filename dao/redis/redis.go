package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		DB:       viper.GetInt("redis.db"),
		Password: viper.GetString("redis.password"),
		PoolSize: viper.GetInt("redis.poolsize"),
	})
	_, err = rdb.Ping().Result()
	zap.L().Fatal("连接redis失败：%v\n", zap.Error(err))
	return
}
