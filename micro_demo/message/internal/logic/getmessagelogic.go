package logic

import (
	"context"

	"micro_demo/message/internal/svc"
	"micro_demo/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *message.MessageReq) (*message.MessageResp, error) {
	// todo: add your logic here and delete this line

	return &message.MessageResp{}, nil
}
