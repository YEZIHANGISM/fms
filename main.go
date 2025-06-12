package main

import (
	"flag"
	"fms/bill"
	"net/http"
	"os"

	"github.com/go-kit/log"
)

func main() {
	// TODO: 放在配置文件中初始化
	listen := flag.String("listen", ":8080", "HTTP listen address")
	flag.Parse()

	// TODO: 使用zap库
	// TODO: 放在全局变量中初始化，这样就不用到处传了
	logger := log.NewLogfmtLogger(os.Stderr)
	// 设置日志的格式，打印更多有用信息
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	svc := bill.NewBillService()
	svc = bill.LoggingService(logger, svc)

	mux := http.NewServeMux()
	// FIXME: 为什么下面的语句不生效？请求的时候会404
	// mux.Handle("/bill/v1/", bill.BuildHandler(svc, logger))
	mux.Handle("/bills", bill.ListBillHandler(svc, logger))

	http.Handle("/", mux)

	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
