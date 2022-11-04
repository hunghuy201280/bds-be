package rebiz

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
)

type ListRealEstateStore interface {
	ListRealEstate(
		ctx context.Context,
		conditions map[string]any,
		filter *remodel.Filter,
		paging *common.Paging,
		moreKeys ...string,

	) ([]remodel.RealEstate, error)

	ListRealEstateImages(
		ctx context.Context,
		reId int,
	) ([]common.Image, error)
}

type listRealEstateBiz struct {
	store ListRealEstateStore
}

func NewListRealEstateBiz(store ListRealEstateStore) *listRealEstateBiz {
	return &listRealEstateBiz{
		store: store,
	}
}

func (biz *listRealEstateBiz) ListRealEstate(
	ctx context.Context,
	filter *remodel.Filter,
	paging *common.Paging) ([]remodel.RealEstate, error) {
	result, err := biz.store.ListRealEstate(ctx, common.JS{"status": 1}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.GetType(remodel.RealEstate{}), err)
	}

	return result, nil

}
