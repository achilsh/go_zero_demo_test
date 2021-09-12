package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"helloworld/internal/logic"
	"helloworld/internal/svc"
	"helloworld/internal/types"
)

func HelloworldHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloworldLogic(r.Context(), ctx)
		resp, err := l.Helloworld(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
