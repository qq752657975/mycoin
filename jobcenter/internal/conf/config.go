package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"jobcenter/internal/database"
	"jobcenter/internal/logic"
)

type Config struct {
	Okx        kline.OkxConfig
	Mongo      database.MongoConfig
	Kafka      database.KafkaConfig
	CacheRedis cache.CacheConf
}
