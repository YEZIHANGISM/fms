package bill

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makelistBillEndpoint(svc BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.ListBill(ctx)
		if err != nil {
			return listBillResponse{v, err.Error()}, nil
		}
		return listBillResponse{v, ""}, nil
	}
}
