package logic

import (
	"context"
	"fmt"

	"multi_micro/api/internal/svc"
	"multi_micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"multi_micro/rpc_one/rpcone"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line
	r, e := l.svcCtx.Adder.Greet(l.ctx, &rpcone.RpcOneReq{})
	if e != nil {
		return nil, fmt.Errorf("call rpc one fail, %v",e)
	}

	_ = r

	return &types.Response{
		Ok: true,
	}, nil
}
