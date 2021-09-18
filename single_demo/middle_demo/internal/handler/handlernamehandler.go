package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"middle_demo/internal/logic"
	"middle_demo/internal/svc"
	"middle_demo/internal/types"
)

func handlerNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHandlerNameLogic(r.Context(), ctx)
		resp, err := l.HandlerName(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
