package main

import (
	"fms/bill"
	"fms/configs"
	"net/http"

	kitzap "github.com/go-kit/kit/log/zap"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	configs.InitFmsConfig()

	// TODO: 放在全局变量中初始化，这样就不用到处传了
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	// go-kit/log/zap，内置了zap库的支持
	logger := kitzap.NewZapSugarLogger(zapLogger, zapcore.DebugLevel)

	svc := bill.NewBillService()
	svc = bill.LoggingService(logger, svc)

	r := mux.NewRouter()
	bill.InitBillRouter(r, svc, logger)

	http.Handle("/", r)

	addr := configs.Config.Basic.FullAddr()
	logger.Log("msg", "HTTP", "addr", addr)
	logger.Log("err", http.ListenAndServe(addr, nil))
}
