package logic

import (
	"context"
	"micro_demo/model"

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

	su := model.Shorturl{
		Shorten: "1111",
		Url:     "http://xxx/yy/ccc",
	}

	_, e := l.svcCtx.Mem.Insert(su)
	if e != nil {
		l.Errorf("insert shorturl fail, %v", e)
		return nil, e
	}

	suOne := "1111"
	r, e := l.svcCtx.Mem.FindOne(suOne)
	if e != nil {
		l.Errorf("find ret, e: %v", e)
		return nil, e
	}
	l.Info("query sql url: %s, short: %s", r.Url, r.Shorten)
	//

	return &message.MessageResp{
		Greet: in.Name + " , this rpc server response",
	}, nil
}
