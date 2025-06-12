package bill

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func ListBillHandler(svc BillService, logger log.Logger) *httptransport.Server {

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		// httptransport.ServerErrorEncoder(encodeError),
	}
	billHandler := httptransport.NewServer(
		makelistBillEndpoint(svc),
		decodelistBillRequest,
		encodeResponse,
		opts...,
	)
	return billHandler
}

func BuildHandler(svc BillService, logger log.Logger) http.Handler {
	// 添加错误日志打印
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		// httptransport.ServerErrorEncoder(encodeError),
	}
	billHandler := httptransport.NewServer(
		makelistBillEndpoint(svc),
		decodelistBillRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/bills", billHandler).Methods("GET")

	return r
}

// type listBillRequest struct{}

type listBillResponse struct {
	Data []Bill `json:"data"`
	Err  string `json:"err,omitempty"`
}

func decodelistBillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
	// var request listBillRequest
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
