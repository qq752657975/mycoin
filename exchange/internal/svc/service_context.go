package svc

import (
	"exchange/internal/config"
	"exchange/internal/consumer"
	"exchange/internal/database"
	"exchange/internal/processor"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/uclient"
	mydb "mycoin-common/msdb"
)

type ServiceContext struct {
	Config      config.Config
	Cache       cache.Cache
	Db          *mydb.MsDB
	MongoClient *database.MongoClient
	MemberRpc   uclient.Member
	MarketRpc   mclient.Market
	AssetRpc    uclient.Asset
	KafkaClient *database.KafkaClient
}

func (sc *ServiceContext) init() {
	factory := processor.NewCoinTradeFactory()
	factory.Init(sc.MarketRpc, sc.KafkaClient, sc.Db)
	kafkaConsumer := consumer.NewKafkaConsumer(sc.KafkaClient, factory, sc.Db)
	kafkaConsumer.Run()
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("market"),
		nil,
		func(o *cache.Options) {})
	kafkaClient := database.NewKafkaClient(c.Kafka)
	client, _ := zrpc.NewClient(c.UCenterRpc)
	newClient, _ := zrpc.NewClient(c.MarketRpc)
	s := &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql),
		MongoClient: database.ConnectMongo(c.Mongo),
		MemberRpc:   uclient.NewMember(client),
		MarketRpc:   mclient.NewMarket(newClient),
		AssetRpc:    uclient.NewAsset(client),
		KafkaClient: kafkaClient,
	}
	s.init()
	return s
}
