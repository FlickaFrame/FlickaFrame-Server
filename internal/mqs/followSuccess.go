package mqs

import (
    "context"
		"fmt"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
)

type FollowSuccess struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewFollowSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *FollowSuccess {
    return &FollowSuccess{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *FollowSuccess) Consume(key, val string) error {
    logx.WithContext(l.ctx).Info("FollowSuccess key :%s , val :%s", key, val)
		fmt.Println("FollowSuccess key :%s , val :%s", key, val)
    return nil
}