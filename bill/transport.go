package bill

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitBillRouter(r *mux.Router, svc BillService) {
	listBillHandler := httptransport.NewServer(
		makelistBillEndpoint(svc),
		decodelistBillRequest,
		encodeResponse,
	)
	createBillHandler := httptransport.NewServer(
		makeCreateBillEndpoint(svc),
		decodeCreateBillRequest,
		encodeResponse,
	)

	sr := r.PathPrefix("/bill/v1/").Subrouter()
	sr.Handle("/bills", listBillHandler).Methods("GET")
	sr.Handle("/bills", createBillHandler).Methods("POST")
}

// type listBillRequest struct{}

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

func decodeCreateBillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateBillRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
