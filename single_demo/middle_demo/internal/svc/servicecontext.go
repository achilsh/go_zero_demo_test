package svc

import (
	"context"
	"github.com/tal-tech/go-zero/rest"
	"middle_demo/internal/config"
	"middle_demo/internal/middleware"
)

type ServiceContext struct {
	Config       config.Config
	AuthMiddle   rest.Middleware //需要手动定义; 需要手动定义,变量名可以和 api文件定义的中间件@server{}中key为： middle对应的value
	AuthMiddle_2 rest.Middleware //需要手动定义; 变量名可以和 api文件定义的中间件@server{}中key为： middle对应的value
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AuthMiddle:   middleware.NewAuthMiddleMiddleware(context.Background(), c.Name).Handle, //需要手动添加
		AuthMiddle_2: middleware.NewAuthMiddle_2Middleware().Handle,                           //需要手动添加
	}
}
