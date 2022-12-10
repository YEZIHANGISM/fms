package logic

import (
	"context"

	"fms/service/fms/api/internal/svc"
	"fms/service/fms/api/internal/types"
	"fms/service/fms/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ListBillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBillLogic {
	return &ListBillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBillLogic) ListBill(req *types.BillReq) (resp *types.BillListReply, err error) {
	conn := sqlx.NewMysql(l.svcCtx.Config.DataSource)
	billModel := model.NewBillModel(conn)

	bills, err := billModel.FindMany(l.ctx)
	if err != nil {
		logx.Errorf("Failed to list bills, error: %s", err.Error())
		return
	}

	resp = &types.BillListReply{
		Bills: make([]types.Bill, 0),
	}

	for _, bill := range bills {
		rep := &types.Bill{
			Id:       bill.Id,
			BillType: bill.BillType,
			Amount:   int(bill.Amount),
			AssetIn:  "",
			AssetOut: "",
			Remark:   "",
		}
		resp.Bills = append(resp.Bills, *rep)
	}

	return
}
