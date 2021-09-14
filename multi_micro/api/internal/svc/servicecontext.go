package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"multi_micro/api/internal/config"
	"multi_micro/rpc_one/rpcone"
	"multi_micro/rpc_two/checker"
)

type ServiceContext struct {
	Config config.Config
	Adder   rpcone.RpcOne        // 对应rpc 下的包和 interace
	Checker checker.Checker      // 对应rpc 下的包和 interace
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder: rpcone.NewRpcOne(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
