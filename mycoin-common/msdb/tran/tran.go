package tran

import mydb "mycoin-common/msdb"

// Transaction 事务的操作 一定跟数据库有关 注入数据库的连接 gorm.db
type Transaction interface {
	Action(func(conn mydb.DbConn) error) error
}
