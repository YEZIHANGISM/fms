package main

import (
	"fms/bill"
	"fms/configs"
	fmslog "fms/log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.InitFmsConfig()

	fmslog.InitLogger()

	svc := bill.NewBillService()
	svc = bill.LoggingService(svc)

	r := mux.NewRouter()
	bill.InitBillRouter(r, svc)
	http.Handle("/", r)

	addr := configs.Config.Basic.FullAddr()
	fmslog.SLogger.Infof("msg: HTTP, addr: %s", addr)
	fmslog.SLogger.Infof("err", http.ListenAndServe(addr, nil))
}
