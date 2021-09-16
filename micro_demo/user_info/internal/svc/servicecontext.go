package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"micro_demo/message/messagesvr"
	"micro_demo/model"
	"micro_demo/user_info/internal/config"
)

type ServiceContext struct {
	Config config.Config
	MessageRpc  messagesvr.MessageSvr //增加 依赖的服务 客户端
	ModelHandle model.BookModel //增加数据库的client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//通过 服务名创建 依赖服务的 客户端
		MessageRpc: messagesvr.NewMessageSvr(zrpc.MustNewClient(c.RpcNode)),
		//通过数据库连接配置来创建数据库的client
		ModelHandle: model.NewBookModel(sqlx.NewMysql(c.DataS)),
	}
}
