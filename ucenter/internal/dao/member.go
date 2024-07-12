package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	mydb "mycoin-common/msdb"
	"mycoin-common/msdb/gorms"
	"ucenter/internal/model"
)

type MemberDao struct {
	comm *gorms.GormConn
}

func (m MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.comm.Session(ctx)
	err = session.
		Model(&model.Member{}).
		Where("mobile_phone=?", phone).
		Limit(1).
		Take(&mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func NewMemberDao(db *mydb.MsDB) *MemberDao {
	return &MemberDao{comm: gorms.New(db.Conn)}
}
