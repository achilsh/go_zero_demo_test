// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package server

import (
	"context"

	"multi_micro/rpc_one/internal/logic"
	"multi_micro/rpc_one/internal/svc"
	"multi_micro/rpc_one/rpc_one"
)

type RpcOneServer struct {
	svcCtx *svc.ServiceContext
}

func NewRpcOneServer(svcCtx *svc.ServiceContext) *RpcOneServer {
	return &RpcOneServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcOneServer) Greet(ctx context.Context, in *rpc_one.RpcOneReq) (*rpc_one.RpcOneResp, error) {
	l := logic.NewGreetLogic(ctx, s.svcCtx)
	return l.Greet(in)
}