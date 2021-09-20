package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/timex"
	"net/http"
	"time"
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

		//配置设置的日志级别： info < error < severe
		// 打印一般普通日志； 只有日志级别(level <= info)才行; 写到 access.log
		logx.Infof("this is test: %s", "yes, is test!!")
		logx.Info("this is test .2....")

		//打印错误日志； 日志级别设置level <= error),写到 error.log
		logx.Error("is error 1...")
		logx.Errorf("is error 2....")

		// 打印慢日志; 日志级别设置(level <= error),写到 slow.log
		nowTm := timex.Now()
		time.Sleep(2 * time.Millisecond)
		durTm := timex.Since(nowTm)
		logx.WithDuration(durTm).Slowf("test slow process")

		// 打印stack日志；只有在日志级别设置为(level <= error), 写到 error.log
		logx.ErrorStack("this stack test....")

		// 打印统计 stat 日志；只有日志级别设置为(level <= info)， 写到 stat.log
		logx.Statf("this stat log.....")

		// 打印 severe 日志； 只有日志级别设置为(level <= severe); 写到 severe.log
		logx.Severef("this is severe log...")

		//打印 tracelog 日志；将 trace, span 写到文件中，至于什么日志文件，将依赖于logx.WithContext().后面的接口
		logx.WithContext(r.Context()).Error("this is trace log")

		// Passthrough to next handler if need
		//根据处理结果要提前结束调用链节点(直接返回 return )，那么后续的业务逻辑将被会被处理.
		//return

		next(w, r)
	}
}
