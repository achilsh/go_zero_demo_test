package logic

import (
	"context"

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

	return &types.Response{
		Shorten: req.Url + ", this is test",
	}, nil
}
