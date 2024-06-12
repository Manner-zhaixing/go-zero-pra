// Code generated by goctl. DO NOT EDIT.
// Source: userRpc.proto

package server

import (
	"context"

	"go-zero-pra/user/userRpc/internal/logic"
	"go-zero-pra/user/userRpc/internal/svc"
	"go-zero-pra/user/userRpc/userRpc"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	userRpc.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRpcServer) Ping(ctx context.Context, in *userRpc.Request) (*userRpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
