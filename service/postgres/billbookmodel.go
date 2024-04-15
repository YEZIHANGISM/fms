package postgres

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BillbookModel = (*customBillbookModel)(nil)

type (
	// BillbookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillbookModel.
	BillbookModel interface {
		billbookModel
	}

	customBillbookModel struct {
		*defaultBillbookModel
	}
)

// NewBillbookModel returns a model for the database table.
func NewBillbookModel(conn sqlx.SqlConn) BillbookModel {
	return &customBillbookModel{
		defaultBillbookModel: newBillbookModel(conn),
	}
}
