package publishing

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gov-cms/app/internal/logic/publishing"
	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"
)

func PublishingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishingListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publishing.NewPublishingLogic(r.Context(), svcCtx)
		resp, err := l.Publishing(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
