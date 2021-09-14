package logic

import (
	"context"

	"multi_micro/rpc_one/internal/svc"
	"multi_micro/rpc_one/rpc_one"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GreetLogic) Greet(in *rpc_one.RpcOneReq) (*rpc_one.RpcOneResp, error) {
	// todo: add your logic here and delete this line

	return &rpc_one.RpcOneResp{}, nil
}
