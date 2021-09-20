package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

//自定义(手动) 一个全局的中间

type GlobalMiddleWare struct {
}

func NewGlobalMiddleWare() *GlobalMiddleWare {
	return &GlobalMiddleWare{}
}

func (m *GlobalMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Error("this is global middleware process......")
		// Passthrough to next handler if need
		//根据处理结果要提前结束调用链节点(直接返回 return )，那么后续的业务逻辑将被会被处理.
		//return

		next(w, r)
	}
}
