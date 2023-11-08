package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteCommentLogic {
	return &FavoriteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteCommentLogic) FavoriteComment(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	resp = &types.FavoriteResp{IsFavorite: true}
	doerId := jwt.GetUidFromCtx(l.ctx)
	// 检查评论是否存在
	_, err = l.svcCtx.CommentModel.FindOneComment(l.ctx, util.MustString2Int64(req.TargetId))
	if err != nil {
		logx.Info(err)
		return nil, code.ErrCommentNoExistsError
	}
	err = l.svcCtx.FavoriteModel.CreateCommentFavorite(l.ctx,
		doerId,
		util.MustString2Int64(req.TargetId),
	)
	if err != nil {
		logx.Info(err)
		return nil, code.DuplicateFavoriteErr
	}
	return
}
