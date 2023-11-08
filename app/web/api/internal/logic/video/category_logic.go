package video

import (
	"context"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

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

func (l *CategoryLogic) Category() (resp *types.CategoryResp, err error) {
	categories, err := l.svcCtx.VideoModel.FindCategories(l.ctx)
	if err != nil {
		logx.Info(err)
		return nil, err
	}
	CategoryList := make([]*types.Category, 0, len(categories))
	for _, v := range categories {
		CategoryList = append(CategoryList, &types.Category{
			ID:   strconv.Itoa(int(v.ID)),
			Name: v.Name,
		})
	}
	resp = &types.CategoryResp{
		CategoryList: CategoryList,
	}
	return
}
