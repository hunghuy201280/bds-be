package rebiz

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
)

type ListRealEstatePostStore interface {
	ListRealEstatePost(
		ctx context.Context,
		conditions map[string]any,
		filter *remodel.Filter,
		paging *common.Paging,
		moreKeys ...string,

	) ([]remodel.RealEstatePost, error)
}

type listRealEstatePostBiz struct {
	store ListRealEstatePostStore
}

func NewListRealEstatePostBiz(store ListRealEstatePostStore) *listRealEstatePostBiz {
	return &listRealEstatePostBiz{
		store: store,
	}
}

func (biz *listRealEstatePostBiz) ListRealEstatePost(
	ctx context.Context,
	filter *remodel.Filter,
	paging *common.Paging) ([]remodel.RealEstatePost, error) {
	result, err := biz.store.ListRealEstatePost(ctx, common.JS{"status": 1}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.GetType(remodel.RealEstatePost{}), err)
	}

	return result, nil

}
