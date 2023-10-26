package feed

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryLogic) Category(req *types.CategoryReq) (resp *types.CategoryResp, err error) {
	categories, err := l.svcCtx.VideoModel.FindCategories(l.ctx)
	if err != nil {
		return nil, err
	}
	CategoryList := make([]*types.Category, 0, len(categories))
	for _, v := range categories {
		CategoryList = append(CategoryList, &types.Category{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	resp = &types.CategoryResp{
		CategoryList: CategoryList,
	}
	return
}
