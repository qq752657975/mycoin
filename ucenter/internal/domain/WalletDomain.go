package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mclient"
	mydb "mycoin-common/msdb"
	"mycoin-common/msdb/tran"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
	marketRpc        mclient.Market
	redisCache       cache.Cache
}

func (d *MemberWalletDomain) FindWalletBySymbol(ctx context.Context, id int64, name string, coin *mclient.Coin) (*model.MemberWalletCoin, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, id, name)
	if err != nil {
		return nil, err
	}
	if mw == nil {
		//新建并存储
		mw, walletCoin := model.NewMemberWallet(id, coin)
		err := d.memberWalletRepo.Save(ctx, mw)
		if err != nil {
			return nil, err
		}
		return walletCoin, nil
	}
	nwc := &model.MemberWalletCoin{}
	_ = copier.Copy(nwc, mw)
	nwc.Coin = coin
	return nwc, nil
}

func (d *MemberWalletDomain) Freeze(ctx context.Context, conn mydb.DbConn, userId int64, money float64, symbol string) error {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, userId, symbol)
	if err != nil {
		return err
	}
	if mw.Balance < money {
		return errors.New("余额不足")
	}
	err = d.memberWalletRepo.UpdateFreeze(ctx, conn, userId, symbol, money)
	if err != nil {
		return err
	}
	return nil
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoin(ctx context.Context, memberId int64, coinName string) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, memberId, coinName)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

func (d *MemberWalletDomain) UpdateWalletCoinAndBase(ctx context.Context, baseWallet *model.MemberWallet, coinWallet *model.MemberWallet) error {
	return d.transaction.Action(func(conn mydb.DbConn) error {
		err := d.memberWalletRepo.UpdateWallet(ctx, conn, baseWallet.Id, baseWallet.Balance, baseWallet.FrozenBalance)
		if err != nil {
			return err
		}
		err = d.memberWalletRepo.UpdateWallet(ctx, conn, coinWallet.Id, coinWallet.Balance, coinWallet.FrozenBalance)
		if err != nil {
			return err
		}
		return nil
	})
}
func (d *MemberWalletDomain) FindByAddress(ctx context.Context, address string) (*model.MemberWallet, error) {
	return d.memberWalletRepo.FindByAddress(ctx, address)
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoinId(ctx context.Context, memberId int64, coinId int64) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinId(ctx, memberId, coinId)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

func NewMemberWalletDomain(db *mydb.MsDB, marketRpc mclient.Market, redisCache cache.Cache) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
		marketRpc:        marketRpc,
		redisCache:       redisCache,
	}
}
