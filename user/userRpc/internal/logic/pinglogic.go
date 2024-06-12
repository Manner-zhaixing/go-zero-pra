package logic

import (
	"context"

	"go-zero-pra/user/userRpc/internal/svc"
	"go-zero-pra/user/userRpc/userRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *userRpc.Request) (*userRpc.Response, error) {
	// todo: add your logic here and delete this line

	return &userRpc.Response{}, nil
}
