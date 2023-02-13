package category

import (
	"gov-cms/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gov-cms/app/internal/logic/category"
	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"
)

func ServiceCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := category.NewServiceCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.ServiceCategoryList(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		result.HttpResult(r, w, resp, err)
	}
}
