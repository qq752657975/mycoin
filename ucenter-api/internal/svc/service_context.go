package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/uclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	URegisterRpc uclient.Register
	ULoginRpc    uclient.Login
	UCMemberRpc  uclient.Member
	UCAssetRpc   uclient.Asset
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		URegisterRpc: uclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
		ULoginRpc:    uclient.NewLogin(zrpc.MustNewClient(c.UCenterRpc)),
		UCAssetRpc:   uclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		UCMemberRpc:  uclient.NewMember(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
