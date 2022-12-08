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

func (l *ListBillLogic) ListBill(req *types.BillReq) (resp *types.BillReply, err error) {
	// todo: add your logic here and delete this line

	return
}
