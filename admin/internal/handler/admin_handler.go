package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iot-platform/admin/internal/logic"
	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"
)

func AdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		req.Name = r.URL.Query().Get("name")
		l := logic.NewAdminLogic(r.Context(), svcCtx)

		resp, err := l.Admin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}

	}
}
