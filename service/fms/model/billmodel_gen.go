// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	billFieldNames          = builder.RawFieldNames(&Bill{})
	billRows                = strings.Join(billFieldNames, ",")
	billRowsExpectAutoSet   = strings.Join(stringx.Remove(billFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	billRowsWithPlaceHolder = strings.Join(stringx.Remove(billFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	billModel interface {
		Insert(ctx context.Context, data *Bill) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Bill, error)
		FindOneById(ctx context.Context, id string) (*Bill, error)
		Update(ctx context.Context, data *Bill) error
		Delete(ctx context.Context, id string) error
	}

	defaultBillModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Bill struct {
		Id         string         `db:"id"`
		CategoryId string         `db:"category_id"`
		BillbookId string         `db:"billbook_id"`
		AssetOutId sql.NullString `db:"asset_out_id"`
		AssetInId  sql.NullString `db:"asset_in_id"`
		DebtId     sql.NullString `db:"debt_id"`
		BillType   string         `db:"bill_type"`
		Amount     int64          `db:"amount"`
		Unit       string         `db:"unit"`
		Remark     sql.NullString `db:"remark"`
		CreatedAt  time.Time      `db:"created_at"`
		UpdatedAt  sql.NullTime   `db:"updated_at"`
		IsDeleted  bool           `db:"is_deleted"`
	}
)

func newBillModel(conn sqlx.SqlConn) *defaultBillModel {
	return &defaultBillModel{
		conn:  conn,
		table: "`bill`",
	}
}

func (m *defaultBillModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBillModel) FindOne(ctx context.Context, id string) (*Bill, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billRows, m.table)
	var resp Bill
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBillModel) FindOneById(ctx context.Context, id string) (*Bill, error) {
	var resp Bill
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBillModel) Insert(ctx context.Context, data *Bill) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, billRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.CategoryId, data.BillbookId, data.AssetOutId, data.AssetInId, data.DebtId, data.BillType, data.Amount, data.Unit, data.Remark, data.IsDeleted)
	return ret, err
}

func (m *defaultBillModel) Update(ctx context.Context, newData *Bill) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, billRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.CategoryId, newData.BillbookId, newData.AssetOutId, newData.AssetInId, newData.DebtId, newData.BillType, newData.Amount, newData.Unit, newData.Remark, newData.IsDeleted, newData.Id)
	return err
}

func (m *defaultBillModel) tableName() string {
	return m.table
}
