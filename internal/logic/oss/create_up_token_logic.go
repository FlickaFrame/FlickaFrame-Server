package oss

import (
	"context"
	"fmt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/deckarep/golang-set"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
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
	var (
		accessKey = l.svcCtx.Config.Oss.AccessKeyId
		secretKey = l.svcCtx.Config.Oss.AccessKeySecret
		bucket    = l.svcCtx.Config.Oss.BucketName
	)
	UploadTypeSet := mapset.NewSet("video", "avatar", "cover") // filter upload type
	if !UploadTypeSet.Contains(req.UploadType) {
		err = fmt.Errorf("upload type error")
		logx.Info(err)
		return
	}
	putPolicy := storage.PutPolicy{
		Scope:           fmt.Sprintf("%s:%s/", bucket, req.UploadType),
		IsPrefixalScope: 1,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","videoDuration":"$(avinfo.video.duration)"}`,
	}
	mac := auth.New(accessKey, secretKey)
	putPolicy.Expires = 3600 //1小时有效期

	upToken := putPolicy.UploadToken(mac)

	resp = &types.CreateUpTokenResp{
		UpToken: upToken,
		Expires: 3600,
	}
	return
}
