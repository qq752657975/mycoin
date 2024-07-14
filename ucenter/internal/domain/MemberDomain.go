package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	mydb "mycoin-common/msdb"
	"mycoin-common/tools"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (c *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及数据库的查询
	mem, err := c.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

func (c *MemberDomain) Register(ctx context.Context,
	username string,
	phone string,
	password string,
	country string,
	promotion string,
	partner string) error {
	mem := model.NewMember()
	_ = tools.Default(mem)
	mem.Id = 0
	//首先处理密码 密码要md5加密，但是md5不安全，所以我们给加上一个salt值
	salt, pwd := tools.Encode(password, nil)
	mem.Salt = salt
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.Username = username
	mem.Country = country
	mem.PromotionCode = promotion
	mem.FillSuperPartner(partner)
	mem.MemberLevel = model.GENERAL
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"
	err := c.MemberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error("插入数据库失败")
		return errors.New("注册失败")
	}
	return nil
}

func (c *MemberDomain) UpdateLoginCount(id int64, incr int) {
	err := c.MemberRepo.UpdateLoginCount(context.Background(), id, incr)
	if err != nil {
		logx.Error(err)
	}
}

func NewMemberDomain(db *mydb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}
