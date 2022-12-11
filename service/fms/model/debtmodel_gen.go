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
	debtFieldNames          = builder.RawFieldNames(&Debt{})
	debtRows                = strings.Join(debtFieldNames, ",")
	debtRowsExpectAutoSet   = strings.Join(stringx.Remove(debtFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	debtRowsWithPlaceHolder = strings.Join(stringx.Remove(debtFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	debtModel interface {
		Insert(ctx context.Context, data *Debt) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Debt, error)
		FindOneById(ctx context.Context, id string) (*Debt, error)
		Update(ctx context.Context, data *Debt) error
		Delete(ctx context.Context, id string) error
	}

	defaultDebtModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Debt struct {
		Id        string       `db:"id"`
		DebtType  string       `db:"debt_type"`
		Object    string       `db:"object"`
		CountInto sql.NullBool `db:"count_into"`
		Ended     sql.NullBool `db:"ended"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		IsDeleted bool         `db:"is_deleted"`
	}
)

func newDebtModel(conn sqlx.SqlConn) *defaultDebtModel {
	return &defaultDebtModel{
		conn:  conn,
		table: "`debt`",
	}
}

func (m *defaultDebtModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDebtModel) FindOne(ctx context.Context, id string) (*Debt, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", debtRows, m.table)
	var resp Debt
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

func (m *defaultDebtModel) FindOneById(ctx context.Context, id string) (*Debt, error) {
	var resp Debt
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", debtRows, m.table)
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

func (m *defaultDebtModel) Insert(ctx context.Context, data *Debt) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, debtRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.DebtType, data.Object, data.CountInto, data.Ended, data.IsDeleted)
	return ret, err
}

func (m *defaultDebtModel) Update(ctx context.Context, newData *Debt) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, debtRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DebtType, newData.Object, newData.CountInto, newData.Ended, newData.IsDeleted, newData.Id)
	return err
}

func (m *defaultDebtModel) tableName() string {
	return m.table
}
