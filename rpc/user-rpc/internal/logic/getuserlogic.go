package logic

import (
	"context"

	"go-zero-pra/rpc/user-rpc/internal/svc"
	"go-zero-pra/rpc/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.UserReq) (*pb.UserRes, error) {
	// todo: add your logic here and delete this line

	return &pb.UserRes{}, nil
}
