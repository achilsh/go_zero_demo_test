package svc

import (
	"context"
	"github.com/tal-tech/go-zero/rest"
	"middle_demo/internal/config"
	"middle_demo/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	AuthMiddle rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AuthMiddle: middleware.NewAuthMiddleMiddleware(context.Background(),c.Name).Handle,
	}
}
