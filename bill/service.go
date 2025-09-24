package bill

import (
	"context"
	"fms/database/models"
)

type BillService interface {
	ListBill(ctx context.Context) ([]Bill, error)
}

type billService struct {
	db models.Bill
}

func NewBillService() BillService {
	return billService{
		db: models.Bill{},
	}
}

func (bs billService) ListBill(ctx context.Context) ([]Bill, error) {
	var result []Bill
	bills, err := bs.db.ListBills(nil)
	if err != nil {
		return nil, err
	}
	for _, b := range bills {
		result = append(result, Bill{ID: b.ID})
	}
	return result, nil
}

type Bill struct {
	ID uint32 `json:"id"`
}
