package rebiz

import (
	"bds-service/modules/realestate/remodel"
	"context"
)

type GetAppProfileStorage interface {
	GetAppProfile(ctx context.Context) (*remodel.AppProfile, error)
}

type getAppProfileBiz struct {
	store GetAppProfileStorage
}

func NewGetAppProfileBiz(storage GetAppProfileStorage) *getAppProfileBiz {
	return &getAppProfileBiz{
		store: storage,
	}
}

func (biz *getAppProfileBiz) GetAppProfile(ctx context.Context) (*remodel.AppProfile, error) {

	store := biz.store
	id, err := store.GetAppProfile(ctx)
	if err != nil {
		return nil, err
	}

	return id, nil

}
