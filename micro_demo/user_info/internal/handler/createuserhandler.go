package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"micro_demo/user_info/internal/logic"
	"micro_demo/user_info/internal/svc"
	"micro_demo/user_info/internal/types"
)

func CreateUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), ctx)
		err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
