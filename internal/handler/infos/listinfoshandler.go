package infos

import (
	"net/http"

	"cpw/go-zero-demo/internal/logic/infos"
	"cpw/go-zero-demo/internal/svc"
	"cpw/go-zero-demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListInfosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListInfosReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := infos.NewListInfosLogic(r.Context(), svcCtx)
		resp, err := l.ListInfos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
