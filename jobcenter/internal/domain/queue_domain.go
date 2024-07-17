package domain

import (
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"log"
)

const KLINE1M = "kline_1m"

type QueueDomain struct {
	cli *database.KafkaClient
}

func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	kline := model.NewKline(data, "1m")
	bytes, _ := json.Marshal(kline)
	msg := database.KafkaData{
		Topic: KLINE1M,
		Data:  bytes,
		Key:   []byte(symbol),
	}
	d.cli.Send(msg)
	log.Println("=================发送数据成功==============")
}

func NewQueueDomain(cli *database.KafkaClient) *QueueDomain {
	return &QueueDomain{
		cli: cli,
	}
}
