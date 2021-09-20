package main

import (
	"flag"
	"fmt"
	"middle_demo/internal/middleware"

	"middle_demo/internal/config"
	"middle_demo/internal/handler"
	"middle_demo/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/template.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// [begin] 增加全局中间件逻辑:
	server.Use(middleware.NewGlobalMiddleWare().Handle) //不是调用Handle, 而是将 Handle 函数付给框架.

	// [end]  增加全局中间件逻辑.

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
