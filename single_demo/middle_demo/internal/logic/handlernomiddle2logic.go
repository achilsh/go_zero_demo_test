package logic

import (
	"context"

	"middle_demo/internal/svc"
	"middle_demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HandlerNoMiddle_2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerNoMiddle_2Logic(ctx context.Context, svcCtx *svc.ServiceContext) HandlerNoMiddle_2Logic {
	return HandlerNoMiddle_2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerNoMiddle_2Logic) HandlerNoMiddle_2(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
