package rebiz

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/common/entitycommon"
	"bds-go-auth-service/modules/realestate/remodel"
	"context"
)

type CreateRealEstateStore interface {
	CreateRealEstate(ctx context.Context, data *remodel.RealEstateCreate) (int, error)
}

type createRealEstateBiz struct {
	store CreateRealEstateStore
}

func NewCreateRealEstateBiz(store CreateRealEstateStore) *createRealEstateBiz {
	return &createRealEstateBiz{store: store}
}

func (biz *createRealEstateBiz) CreateRealEstate(ctx context.Context, data *remodel.RealEstateCreate) (int, error) {
	if err := data.Validate(); err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrCannotCreateEntity(data.TableName(), err)
	}
	id, err := biz.store.CreateRealEstate(ctx, data)
	if err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrCannotCreateEntity(data.TableName(), err)
	}
	return id, nil
}
