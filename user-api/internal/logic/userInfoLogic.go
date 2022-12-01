package logic

import (
	"context"
	"go-zero-rpc/user-rpc/pb"

	"go-zero-rpc/user-api/internal/svc"
	"go-zero-rpc/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	return &types.UserInfoResp{
		UserId:   user.Id,
		Nickname: user.Nickname,
	}, nil
}
