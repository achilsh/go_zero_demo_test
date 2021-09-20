package handler

import (
	"net/http"

	"middle_demo/internal/logic"
	"middle_demo/internal/svc"
	"middle_demo/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func handlerNomiddleNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHandlerNomiddleNameLogic(r.Context(), ctx)
		resp, err := l.HandlerNomiddleName(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
