package mqs

import (
    "context"
	"encoding/json"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	notice_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/notice"
)

type ActionSuccess struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewActionSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *ActionSuccess {
    return &ActionSuccess{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ActionSuccess) Consume(key, val string) (err error) {
    logx.WithContext(l.ctx).Info("ActionSuccess: ", key, val)
    var notice notice_model.Notice
    err = json.Unmarshal([]byte(val), &notice)
    l.svcCtx.NoticeModel.Insert(l.ctx, &notice)
    return nil
}