package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add zrpc.RpcClientConf // 新增rpc 节点配置，名字和etc/book-api.yaml 文件中定义一致，这样业务服务就可以调用依赖的服务
	Check zrpc.RpcClientConf // 新增rpc 节点配置，名字和etc/book-api.yaml 文件中定义一致，这样业务服务就可以调用依赖的服务
}
