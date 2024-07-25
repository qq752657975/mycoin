package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/uclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc uclient.Register
	UCLoginRpc    uclient.Login
	UCAssetRpc    uclient.Asset
	UCMemberRpc   uclient.Member
	UCWithdrawRpc uclient.Withdraw
	MarketRpc     mclient.Market
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: uclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
		UCLoginRpc:    uclient.NewLogin(zrpc.MustNewClient(c.UCenterRpc)),
		UCAssetRpc:    uclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		UCMemberRpc:   uclient.NewMember(zrpc.MustNewClient(c.UCenterRpc)),
		UCWithdrawRpc: uclient.NewWithdraw(zrpc.MustNewClient(c.UCenterRpc)),
		MarketRpc:     mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
	}
}
