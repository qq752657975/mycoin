package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetLogic struct {
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	logger             logx.Logger
	captchaDomain      *domain.CaptchaDomain
	memberDomain       *domain.MemberDomain
	memberWalletDomain *domain.MemberWalletDomain
}

func (l *AssetLogic) FindWalletBySymbol(req *asset.AssetReq) (*asset.MemberWallet, error) {
	//通过market rpc 进行coin表的查询 coin信息
	//通过钱包 查询对应币的钱包信息  coin_id  user_id 查询用户的钱包信息 组装信息
	coinInfo, err := l.svcCtx.MarketRpc.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: req.CoinName,
	})
	if err != nil {
		return nil, err
	}
	memberWalletCoin, err := l.memberWalletDomain.FindWalletBySymbol(l.ctx, req.UserId, req.CoinName, coinInfo)
	if err != nil {
		return nil, err
	}
	resp := &asset.MemberWallet{}
	_ = copier.Copy(resp, memberWalletCoin)
	return resp, nil
}

func (l *AssetLogic) FindWallet(req *asset.AssetReq) (*asset.MemberWalletList, error) {
	//根据用户id查询用户的钱包 循环钱包信息 根据币种 查询币种详情
	memberWalletCoins, err := l.memberWalletDomain.FindWallet(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberWallet
	copier.Copy(&list, memberWalletCoins)
	return &asset.MemberWalletList{
		List: list,
	}, nil
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		logger:             logx.WithContext(ctx),
		captchaDomain:      domain.NewCaptchaDomain(),
		memberDomain:       domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain: domain.NewMemberWalletDomain(svcCtx.Db, svcCtx.MarketRpc, svcCtx.Cache),
	}
}
