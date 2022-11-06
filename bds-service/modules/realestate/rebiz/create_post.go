package rebiz

import (
	"bds-service/common"
	"bds-service/common/entitycommon"
	"bds-service/modules/realestate/remodel"
	"context"
	"errors"
	"github.com/gookit/goutil"
)

type CreatePostStorage interface {
	CreatePost(ctx context.Context, post *remodel.RealEstatePostCreate) (int, error)
}

type createPostBiz struct {
	store CreatePostStorage
}

func NewCreatePostBiz(storage CreatePostStorage) *createPostBiz {
	return &createPostBiz{
		store: storage,
	}
}

func (biz *createPostBiz) CreatePost(ctx context.Context, post *remodel.RealEstatePostCreate) (int, error) {
	if goutil.IsEmpty(post) {
		return entitycommon.ERROR_ENTITY_ID, common.ErrInvalidRequest(errors.New("data is empty"))
	}

	if err := post.Validate(); err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrInvalidRequest(err)
	}
	store := biz.store
	id, err := store.CreatePost(ctx, post)
	if err != nil {
		return entitycommon.ERROR_ENTITY_ID, err
	}

	return id, nil

}
