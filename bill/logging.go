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

func (mw loggingMiddleware) ListBills(ctx context.Context) (data []Bill, err error) {
	defer func(begin time.Time) {
		mw.logger.Infof(
			"message: listBills, "+
				"data: %v, "+
				"err: %v, "+
				"took: %v",
			data,
			err,
			time.Since(begin),
		)
	}(time.Now())
	data, err = mw.next.ListBills(ctx)
	return
}

func (mw loggingMiddleware) CreateBill(ctx context.Context, bill CreateBillRequest) (err error) {
	defer func(begin time.Time) {
		mw.logger.Infof(
			"message: createBill, "+
				"bill: %v, "+
				"err: %v, "+
				"took: %v",
			bill,
			err,
			time.Since(begin),
		)
	}(time.Now())
	err = mw.next.CreateBill(ctx, bill)
	return
}
