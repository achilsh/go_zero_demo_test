package logic

import (
	"context"

	"multi_micro/rpc_two/internal/svc"
	"multi_micro/rpc_two/rpc_two"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *rpc_two.CheckReq) (*rpc_two.CheckResp, error) {
	// todo: add your logic here and delete this line

	return &rpc_two.CheckResp{}, nil
}
