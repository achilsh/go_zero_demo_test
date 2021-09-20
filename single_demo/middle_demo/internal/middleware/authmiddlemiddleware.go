package middleware

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type AuthMiddleMiddleware struct {
	ResData string //增加业务数据
	logx.Logger
}

func NewAuthMiddleMiddleware(ctx context.Context, data string) *AuthMiddleMiddleware {
	return &AuthMiddleMiddleware{
		Logger:  logx.WithContext(ctx),
		ResData: data, //创建中间件数据
	}
}

func (m *AuthMiddleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		m.Infof("this middleware, data: %s", m.ResData) //增加中间件消息处理
		// Passthrough to next handler if need
		next(w, r)
	}
}
