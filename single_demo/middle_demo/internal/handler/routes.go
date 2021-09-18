// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"middle_demo/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddle},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/users/id/:userId",
					Handler: handlerNameHandler(serverCtx),
				},
			}...,
		),
	)
}
