package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"multi_micro/api/internal/logic"
	"multi_micro/api/internal/svc"
	"multi_micro/api/internal/types"
)

func AddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
