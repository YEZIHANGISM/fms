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
	billbookFieldNames          = builder.RawFieldNames(&Billbook{})
	billbookRows                = strings.Join(billbookFieldNames, ",")
	billbookRowsExpectAutoSet   = strings.Join(stringx.Remove(billbookFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	billbookRowsWithPlaceHolder = strings.Join(stringx.Remove(billbookFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	billbookModel interface {
		Insert(ctx context.Context, data *Billbook) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Billbook, error)
		FindOneById(ctx context.Context, id string) (*Billbook, error)
		Update(ctx context.Context, data *Billbook) error
		Delete(ctx context.Context, id string) error
	}

	defaultBillbookModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Billbook struct {
		Id        string       `db:"id"`
		Name      string       `db:"name"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		IsDeleted bool         `db:"is_deleted"`
	}
)

func newBillbookModel(conn sqlx.SqlConn) *defaultBillbookModel {
	return &defaultBillbookModel{
		conn:  conn,
		table: "`billbook`",
	}
}

func (m *defaultBillbookModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBillbookModel) FindOne(ctx context.Context, id string) (*Billbook, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billbookRows, m.table)
	var resp Billbook
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

func (m *defaultBillbookModel) FindOneById(ctx context.Context, id string) (*Billbook, error) {
	var resp Billbook
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billbookRows, m.table)
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

func (m *defaultBillbookModel) Insert(ctx context.Context, data *Billbook) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, billbookRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.IsDeleted)
	return ret, err
}

func (m *defaultBillbookModel) Update(ctx context.Context, newData *Billbook) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, billbookRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.IsDeleted, newData.Id)
	return err
}

func (m *defaultBillbookModel) tableName() string {
	return m.table
}