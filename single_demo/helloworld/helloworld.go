package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"helloworld/internal/config"
	"helloworld/internal/handler"
	"helloworld/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/helloworld-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 根据配置文件项 来初始化 log
	logx.MustSetup(c.Log)
	// defer logx.Close() //不需要关闭，因为在 server.Stop()中关闭日志句柄
	// logx.SetLevel(logx.SevereLevel) //该接口可修改日志的级别， 目前级别有
	//	// 1. InfoLevel logs everything
	//	InfoLevel = iota
	//	// 2. ErrorLevel includes errors, slows, stacks
	//	ErrorLevel
	//	// 3. SevereLevel only log severe messages
	//	SevereLevel

	logx.Info("init main process...")

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
