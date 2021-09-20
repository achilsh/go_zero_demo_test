package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/timex"
	"micro_demo/message/messagesvr"
	"micro_demo/model"
	"time"

	"micro_demo/user_info/internal/svc"
	"micro_demo/user_info/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

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
	logx.WithContext(l.ctx).Error("this is trace log")

	//insert db op ...
	b := model.Book{
		Book:  "this test book: 1",
		Price: 12312,
	}

	if _, e := l.svcCtx.ModelHandle.Insert(b); e != nil {
		l.Errorf("insert tab: %s fail, err: %v", l.svcCtx.Config.Table, e)
		return nil, e
	}
	//find db op ...

	BookVal := "this test book: 0"
	retb, e := l.svcCtx.ModelHandle.FindOne(BookVal)
	if e != nil {
		l.Errorf("find by sql fail, err: %v", e)
		return nil, e
	}
	if retb == nil {
		l.Errorf("find by sql ret is nil")
		return nil, errors.New("ret is nil")
	}
	l.Info("find book: %v", retb)

	// 调用依赖服务： 通过上下文获取依赖服务的client， 调用对应接口
	ret, e := l.svcCtx.MessageRpc.GetMessage(l.ctx, &messagesvr.MessageReq{Name: "req data for rpc"})
	if e != nil {
		l.Errorf("call GetMessage() fail, %v", e)
		return nil, fmt.Errorf("call fail, %v", e)
	}

	l.Info("call rpc succ.............")
	return &types.Response{
		Shorten: req.Url + ", rpc rsp: " + ret.Greet,
	}, nil
}
