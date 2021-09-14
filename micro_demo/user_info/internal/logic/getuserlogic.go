package logic

import (
	"context"
	"fmt"
	"micro_demo/message/messagesvr"

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
