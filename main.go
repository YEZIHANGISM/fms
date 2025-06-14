package main

import (
	"flag"
	"fms/bill"
	"net/http"

	kitzap "github.com/go-kit/kit/log/zap"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// TODO: 放在配置文件中初始化
	listen := flag.String("listen", ":8080", "HTTP listen address")
	flag.Parse()

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

	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
