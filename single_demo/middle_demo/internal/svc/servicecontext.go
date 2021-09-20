package svc

import (
	"context"
	"github.com/tal-tech/go-zero/rest"
	"middle_demo/internal/config"
	"middle_demo/internal/middleware"
)

type ServiceContext struct {
	Config       config.Config
	AuthMiddle   rest.Middleware //需要手动定义
	Middle2      rest.Middleware //需要手动定义
	AuthMiddle_2 rest.Middleware //需要手动定义
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AuthMiddle:   middleware.NewAuthMiddleMiddleware(context.Background(), c.Name).Handle, //需要手动添加
		Middle2:      middleware.NewAuthMiddle_2Middleware().Handle,                           //需要手动添加
		AuthMiddle_2: middleware.NewAuthMiddle_2Middleware().Handle,                           //需要手动添加
	}
}
