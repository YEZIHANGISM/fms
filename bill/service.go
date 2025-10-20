package bill

import (
	"context"
	"fms/database/models"
	fmslog "fms/log"
	"time"
)

type BillService interface {
	ListBills(ctx context.Context) ([]Bill, error)
	CreateBill(ctx context.Context, req CreateBillRequest) error
}

type billService struct {
	db models.Bill
}

func NewBillService() BillService {
	return billService{
		db: models.Bill{},
	}
}

func (bs billService) ListBills(ctx context.Context) ([]Bill, error) {
	var result []Bill
	bills, err := bs.db.ListBills(nil)
	if err != nil {
		return nil, err
	}
	for _, b := range bills {
		result = append(result, Bill{
			ID:         b.ID,
			BillBookID: b.BillbookID,
			CategoryID: b.CategoryID,
			AssetOutID: b.AssetOutID,
			AssetInID:  b.AssetInID,
			DebtID:     b.DebtID,
			BillDate:   time.Duration(b.BillDate.UnixMilli()),
			Type:       b.Type,
			Amount:     b.Amount,
			Remark:     b.Remark,
		})
	}
	return result, nil
}

func (bs billService) CreateBill(ctx context.Context, req CreateBillRequest) error {
	fmslog.SLogger.Infof("Creating bill, req: %#v", req)
	record := models.Bill{
		BillbookID: req.BillBookID,
		CategoryID: req.CategoryID,
		AssetOutID: req.AssetOutID,
		AssetInID:  req.AssetInID,
		DebtID:     req.DebtID,
		BillDate:   time.UnixMilli(int64(req.BillDate)),
		Type:       req.Type,
		Amount:     req.Amount,
		Remark:     req.Remark,
	}
	return bs.db.CreateBill(ctx, &record)
}
