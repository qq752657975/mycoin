package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	mydb "mycoin-common/msdb"
	"ucenter/internal/config"
	"ucenter/internal/database"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *mydb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mycoin"),
		nil,
		func(o *cache.Options) {
		})

	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     database.ConnMysql(c.Mysql.DataSource),
	}
}
