package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillModel = (*customBillModel)(nil)

type (
	// BillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillModel.
	BillModel interface {
		billModel
		FindMany(ctx context.Context) ([]*Bill, error)
	}

	customBillModel struct {
		*defaultBillModel
	}
)

// NewBillModel returns a model for the database table.
func NewBillModel(conn sqlx.SqlConn) BillModel {
	return &customBillModel{
		defaultBillModel: newBillModel(conn),
	}
}

func (m *customBillModel) FindMany(ctx context.Context) (bills []*Bill, err error) {
	querySQL := `select ` + billRows + ` from ` + m.table
	err = m.conn.QueryRows(&bills, querySQL)
	if err != nil {
		logx.Errorf("bill.FindMany error, err=%s", err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return
	}
	return
}
