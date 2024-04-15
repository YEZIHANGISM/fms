package model

import (
	"context"
	"fms/service/fms/database"
	"time"

	"go.etcd.io/bbolt"
)

type Bill struct {
	Id         string
	CategoryId string
	BillbookId string
	AssetOutId *string
	AssetInId  *string
	DebtId     *string
	BillType   string
	Amount     int64
	Unit       string
	Remark     *string
	BillTime   time.Time
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	IsDeleted  bool
}

type BillCond struct {
	BookId    string
	StartedAt *int64
	EndedAt   *int64
	PageNum   int64
	PageSize  int64
}

func ListBill(ctx context.Context, condition *BillCond) ([]*Bill, error) {
	db, err := database.Init()
	if err != nil {
		
	}
}