package bill

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   BillService
}

func LoggingService(logger log.Logger, svc BillService) BillService {
	return loggingMiddleware{
		next:   svc,
		logger: logger,
	}
}

func (mw loggingMiddleware) ListBill(ctx context.Context) (data []Bill, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"data", data,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	data, err = mw.next.ListBill(ctx)
	return
}
