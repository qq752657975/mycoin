package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/uclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	URegisterRpc uclient.Register
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		URegisterRpc: uclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
