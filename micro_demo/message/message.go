package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"micro_demo/message/internal/config"
	"micro_demo/message/internal/server"
	"micro_demo/message/internal/svc"
	"micro_demo/message/message"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

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
	srv := server.NewMessageSvrServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		message.RegisterMessageSvrServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
