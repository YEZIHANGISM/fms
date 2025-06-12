package bill

import (
	"context"
)

type BillService interface {
	ListBill(ctx context.Context) ([]Bill, error)
}

type billService struct{}

func NewBillService() BillService {
	return billService{}
}

type Bill struct {
	ID int `json:"id"`
}

func (billService) ListBill(ctx context.Context) ([]Bill, error) {
	// TODO: 业务逻辑编写
	return []Bill{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}, nil
}
