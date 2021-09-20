package logic

import (
	"context"

	"middle_demo/internal/svc"
	"middle_demo/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HandlerNomiddleNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerNomiddleNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) HandlerNomiddleNameLogic {
	return HandlerNomiddleNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerNomiddleNameLogic) HandlerNomiddleName(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
