package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingLogic) Following(req *types.FollowingReq) (resp *types.FollowingResp, err error) {
	//doerId := jwt.GetUidFromCtx(l.ctx)
	//videos, err := l.svcCtx.VideoModel.FollowingUserVideo(l.ctx, doerId, orm.ListOptions{
	//	PageSize: req.PageSize,
	//	Page:     req.Page,
	//	ListAll:  req.ListAll,
	//})
	//if err != nil {
	//	logx.Debug(err)
	//	return
	//}
	//resp = &types.FollowingResp{}
	//for _, v := range videos { // TODO: 优化
	//feedItem := &types.Video{}
	//author := l.svcCtx.UserModel.MustFindOne(l.ctx, v.AuthorID)
	//copier.Copy(feedItem, v)
	//feedItem.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
	//feedItem.PlayUrl = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, v.PlayUrl)
	//copier.Copy(&feedItem.Author, author)
	//feedItem.Author.UserID = author.ID
	//feedItem.Author.Avatar = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, author.AvatarUrl)
	//resp.VideoList = append(resp.VideoList, feedItem)
	//feedItem.IsFollow = l.svcCtx.UserModel.IsFollowing(l.ctx, doerId, v.AuthorID)
	//feedItem.IsFav, _ = l.svcCtx.FavoriteModel.IsFavoriteVideo(l.ctx, doerId, v.ID)
	//}
	return
}
