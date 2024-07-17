package kline

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"log"
	"mycoin-common/tools"
	"strings"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
	ch          cache.Cache
}

func NewKline(c OkxConfig, client *database.MongoClient, kafkaClient *database.KafkaClient, cache2 cache.Cache) *Kline {
	return &Kline{
		wg:          sync.WaitGroup{},
		c:           c,
		klineDomain: domain.NewKlineDomain(client),
		queueDomain: domain.NewQueueDomain(kafkaClient),
		ch:          cache2,
	}
}
func (k *Kline) Do(period string) {
	log.Println("============启动k线数据拉取==============")
	k.wg.Add(2)
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
	log.Println("===============k线数据拉取结束===============")
}

func (k *Kline) getKlineData(instId string, symbol string, period string) {
	//发起http请求 获取数据
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	logx.Info("==================执行存储mongo====================")
	if result.Code == "0" {
		//代表成功
		k.klineDomain.SaveBatch(result.Data, symbol, period)
		if "1m" == period {
			//把这个最新的数据result.Data[0] 推送到market服务，推送到前端页面，实时进行变化
			//->kafka->market kafka消费者进行数据消费-> 通过websocket通道发送给前端 ->前端更新数据
			if len(result.Data) > 0 {
				data := result.Data[0]
				k.queueDomain.Send1mKline(data, symbol)
				//放入redis 将其最新的价格
				key := strings.ReplaceAll(instId, "-", "::")
				k.ch.Set(key+"::RATE", data[4])
			}
		}
	}
	k.wg.Done()
	logx.Info("==================End====================")
}
