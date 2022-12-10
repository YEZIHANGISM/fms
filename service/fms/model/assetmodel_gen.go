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
	assetFieldNames          = builder.RawFieldNames(&Asset{})
	assetRows                = strings.Join(assetFieldNames, ",")
	assetRowsExpectAutoSet   = strings.Join(stringx.Remove(assetFieldNames, "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	assetRowsWithPlaceHolder = strings.Join(stringx.Remove(assetFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"
)

type (
	assetModel interface {
		Insert(ctx context.Context, data *Asset) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Asset, error)
		FindOneById(ctx context.Context, id string) (*Asset, error)
		Update(ctx context.Context, data *Asset) error
		Delete(ctx context.Context, id string) error
	}

	defaultAssetModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Asset struct {
		Id          string       `db:"id"`
		Name        string       `db:"name"`
		AssetCateId string       `db:"asset_cate_id"`
		Hide        sql.NullBool `db:"hide"`
		CountInto   sql.NullBool `db:"count_into"`
		Balance     int64        `db:"balance"`
		Unit        string       `db:"unit"`
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   sql.NullTime `db:"updated_at"`
		IsDeleted   bool         `db:"is_deleted"`
	}
)

func newAssetModel(conn sqlx.SqlConn) *defaultAssetModel {
	return &defaultAssetModel{
		conn:  conn,
		table: "`asset`",
	}
}

func (m *defaultAssetModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssetModel) FindOne(ctx context.Context, id string) (*Asset, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", assetRows, m.table)
	var resp Asset
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

func (m *defaultAssetModel) FindOneById(ctx context.Context, id string) (*Asset, error) {
	var resp Asset
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", assetRows, m.table)
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

func (m *defaultAssetModel) Insert(ctx context.Context, data *Asset) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, assetRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.AssetCateId, data.Hide, data.CountInto, data.Balance, data.Unit, data.IsDeleted)
	return ret, err
}

func (m *defaultAssetModel) Update(ctx context.Context, newData *Asset) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, assetRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.AssetCateId, newData.Hide, newData.CountInto, newData.Balance, newData.Unit, newData.IsDeleted, newData.Id)
	return err
}

func (m *defaultAssetModel) tableName() string {
	return m.table
}
