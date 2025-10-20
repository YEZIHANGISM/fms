package bill

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type listBillResponse struct {
	Data []Bill `json:"data"`
	Err  string `json:"err,omitempty"`
}

type Bill struct {
	ID         int           `json:"id"`
	BillBookID int           `json:"billbook_id"`
	CategoryID int           `json:"category_id"`
	AssetOutID int           `json:"asset_out_id"`
	AssetInID  int           `json:"asset_in_id"`
	DebtID     int           `json:"debt_id"`
	BillDate   time.Duration `json:"bill_date"`
	Type       int           `json:"type"`
	Amount     float32       `json:"amount"`
	Remark     string        `json:"remark"`
}

func makelistBillEndpoint(svc BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.ListBills(ctx)
		if err != nil {
			return listBillResponse{v, err.Error()}, nil
		}
		return listBillResponse{v, ""}, nil
	}
}

type CreateBillRequest struct {
	BillBookID int           `json:"billbook_id"`
	CategoryID int           `json:"category_id"`
	AssetOutID int           `json:"asset_out_id"`
	AssetInID  int           `json:"asset_in_id"`
	DebtID     int           `json:"debt_id"`
	BillDate   time.Duration `json:"bill_date"`
	Type       int           `json:"type"`
	Amount     float32       `json:"amount"`
	Remark     string        `json:"remark"`
}

type CreateBillResponse struct {
	Err error `json:"error,omitempty"`
}

func makeCreateBillEndpoint(svc BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBillRequest)
		if err := svc.CreateBill(ctx, req); err != nil {
			return CreateBillResponse{err}, nil
		}
		return CreateBillResponse{nil}, nil
	}
}
