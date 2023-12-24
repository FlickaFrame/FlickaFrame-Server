package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUpTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpTokenLogic {
	return &CreateUpTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUpTokenLogic) CreateUpToken(req *types.CreateUpTokenReq) (resp *types.CreateUpTokenResp, err error) {
	token, err := l.svcCtx.OssRpc.CreatUpToken(l.ctx, &oss.CreateUpTokenRequest{
		UploadType: req.UploadType,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.CreateUpTokenResp{
		UpToken: token.Token,
	}
	return
}
