package logic

import (
	"context"

	"helloworld/internal/svc"
	"helloworld/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelloworldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloworldLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelloworldLogic {
	return HelloworldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloworldLogic) Helloworld(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{Message: "this single hello world response"}, nil
}
