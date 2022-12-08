package handler

import (
	"net/http"

	"fms/service/fms/api/internal/logic"
	"fms/service/fms/api/internal/svc"
	"fms/service/fms/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listBillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BillReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListBillLogic(r.Context(), svcCtx)
		resp, err := l.ListBill(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
