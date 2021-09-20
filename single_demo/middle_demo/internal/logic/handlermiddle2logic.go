package logic

import (
	"context"

	"middle_demo/internal/svc"
	"middle_demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HandlerMiddle_2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerMiddle_2Logic(ctx context.Context, svcCtx *svc.ServiceContext) HandlerMiddle_2Logic {
	return HandlerMiddle_2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerMiddle_2Logic) HandlerMiddle_2(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
