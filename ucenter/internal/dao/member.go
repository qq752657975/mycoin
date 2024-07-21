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
	conn *gorms.GormConn
}

func (m *MemberDao) FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("id=?", memberId).Take(&mem).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (m *MemberDao) UpdateLoginCount(ctx context.Context, id int64, incr int) error {
	session := m.conn.Session(ctx)
	err := session.Exec("update member set login_count=login_count+? where id = ?", incr, id).Error
	return err
}

func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
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
	return &MemberDao{conn: gorms.New(db.Conn)}
}
