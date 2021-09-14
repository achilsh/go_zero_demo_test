package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	//添加 rpc client的配置, 变量名就是配置项名
	RpcNode zrpc.RpcClientConf
}
