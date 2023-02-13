package ping

import (
	"gov-cms/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gov-cms/app/internal/logic/ping"
	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := ping.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		result.HttpResult(r, w, resp, err)
	}
}
