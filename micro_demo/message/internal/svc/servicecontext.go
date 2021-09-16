package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"micro_demo/message/internal/config"
	"micro_demo/model"
)

type ServiceContext struct {
	Config config.Config
	Mem    model.ShorturlModel //新增数据模型1,
	Other  model.BookModel     //新增数据模型2, 涉及到多个表
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mem:    model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), //根据数据库配置创建model和cache，此处cache和db关联起来
		Other:  model.NewBookModel(sqlx.NewMysql(c.DataSource)),
	}
}
