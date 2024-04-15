package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillTypeModel = (*customBillTypeModel)(nil)

type (
	// BillTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillTypeModel.
	BillTypeModel interface {
		billTypeModel
	}

	customBillTypeModel struct {
		*defaultBillTypeModel
	}
)

// NewBillTypeModel returns a model for the database table.
func NewBillTypeModel(conn sqlx.SqlConn) BillTypeModel {
	return &customBillTypeModel{
		defaultBillTypeModel: newBillTypeModel(conn),
	}
}
