package ucenter

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-api/api/code/internal/logic/ucenter"
	"go-zero-api/api/code/internal/svc"
	"go-zero-api/api/code/internal/types"
)

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ucenter.NewDownloadFileLogic(r.Context(), svcCtx, w)
		err := l.DownloadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
