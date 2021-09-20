package handler

import (
	"net/http"

	"middle_demo/internal/logic"
	"middle_demo/internal/svc"
	"middle_demo/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func handlerNoMiddle_2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHandlerNoMiddle_2Logic(r.Context(), ctx)
		resp, err := l.HandlerNoMiddle_2(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
