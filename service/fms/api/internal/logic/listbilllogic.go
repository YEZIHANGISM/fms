package logic

import (
	"context"

	"fms/service/fms/api/internal/svc"
	"fms/service/fms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	resp = &types.BillListReply{
		Bills: make([]types.Bill, 0),
	}

	rep := &types.Bill{
		Id:       1,
		BillType: "income",
		Amount:   -123.12,
		AssetIn:  "",
		AssetOut: "cash",
		Remark:   "test",
	}
	resp.Bills = append(resp.Bills, *rep)

	return
}
