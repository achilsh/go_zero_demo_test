package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"micro_demo/message/messagesvr"
	"micro_demo/user_info/internal/config"
)

type ServiceContext struct {
	Config config.Config
	MessageRpc  messagesvr.MessageSvr //增加 依赖的服务 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//通过 服务名创建 依赖服务的 客户端
		MessageRpc: messagesvr.NewMessageSvr(zrpc.MustNewClient(c.RpcNode)),
	}
}
