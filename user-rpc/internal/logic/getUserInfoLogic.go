package logic

import (
	"context"
	"fmt"
	"go-zero-rpc/user-rpc/internal/svc"
	"go-zero-rpc/user-rpc/pb"
	"google.golang.org/grpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Printf("tmp: %+v \n", tmp)
	}
	return nil, nil
}
