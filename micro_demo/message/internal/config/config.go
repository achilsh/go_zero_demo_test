package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string          //db 配置,可以配置多个db数据源,名字和config 文件的key保持一致
	Table      string          //某一个表名，当然可以配合多个表名， 名字和config文件的key保持一致; 实际查询不使用该配置。
	Cache      cache.CacheConf //cache 配置
}
