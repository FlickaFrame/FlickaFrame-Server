package logic

import (
	"context"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatUpTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatUpTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatUpTokenLogic {
	return &CreatUpTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatUpTokenLogic) CreatUpToken(in *oss_service.CreateUpTokenRequest) (*oss_service.CreateUpTokenResponse, error) {
	var (
		accessKey = l.svcCtx.Config.Oss.AccessKeyId
		secretKey = l.svcCtx.Config.Oss.AccessKeySecret
		bucket    = l.svcCtx.Config.Oss.BucketName
	)
	UploadTypeSet := mapset.NewSet("video", "avatar", "cover") // filter upload type
	if !UploadTypeSet.Contains(in.UploadType) {
		err := fmt.Errorf("upload type error")
		logx.Info(err)
		return nil, err
	}
	putPolicy := storage.PutPolicy{
		Scope:           fmt.Sprintf("%s:%s/", bucket, in.UploadType),
		IsPrefixalScope: 1,
		ReturnBody:      `{"key":"$(key)","hash":"$(etag)","videoHeight":"$(avinfo.video.height)","videoWidth":"$(avinfo.video.width)","videoDuration":"$(avinfo.video.duration)"}`,
	}
	mac := auth.New(accessKey, secretKey)
	putPolicy.Expires = 3600 //1小时有效期

	upToken := putPolicy.UploadToken(mac)

	resp := &oss_service.CreateUpTokenResponse{
		Token: upToken,
	}
	return resp, nil
}
