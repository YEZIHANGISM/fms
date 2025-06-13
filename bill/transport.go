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

func InitBillRouter(r *mux.Router, svc BillService, logger log.Logger) {
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

	sr := r.PathPrefix("/bill/v1/").Subrouter()
	sr.Handle("/bills", billHandler).Methods("GET")
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
