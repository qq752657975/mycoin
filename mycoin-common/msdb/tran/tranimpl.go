package tran

import (
	"gorm.io/gorm"
	mydb "mycoin-common/msdb"
	"mycoin-common/msdb/gorms"
)

type TransactionImpl struct {
	conn mydb.DbConn
}

func (t *TransactionImpl) Action(f func(conn mydb.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction(db *gorm.DB) *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.New(db),
	}
}
