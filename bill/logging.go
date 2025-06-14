package bill

import (
	"context"
	fmslog "fms/log"
	"time"

	"go.uber.org/zap"
)

type loggingMiddleware struct {
	logger zap.SugaredLogger
	next   BillService
}

func LoggingService(svc BillService) BillService {
	return loggingMiddleware{
		next:   svc,
		logger: *fmslog.SLogger,
	}
}

func (mw loggingMiddleware) ListBill(ctx context.Context) (data []Bill, err error) {
	defer func(begin time.Time) {
		mw.logger.Infof(
			"message: listBill, "+
				"data: %v, "+
				"err: %v, "+
				"took: %v",
			data,
			err,
			time.Since(begin),
		)
	}(time.Now())
	data, err = mw.next.ListBill(ctx)
	return
}
