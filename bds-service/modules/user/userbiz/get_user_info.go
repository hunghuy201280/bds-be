package userbiz

import (
	"bds-service/common"
	"bds-service/modules/user/usermodel"
	"context"
)

type GetUserInfoStore interface {
	FindOneUser(ctx context.Context, data *usermodel.UserFind) (*usermodel.UserLite, error)
}

type getUserInfoBiz struct {
	store GetUserInfoStore
}

func NewGetUserInfoBiz(store GetUserInfoStore) *getUserInfoBiz {
	return &getUserInfoBiz{
		store: store,
	}
}

func (biz *getUserInfoBiz) GetUserInfo(ctx context.Context, userId int) (*usermodel.UserLite, error) {
	store := biz.store
	filter := usermodel.UserFind{}
	filter.Id = userId
	user, err := store.FindOneUser(ctx, &filter)
	if err != nil {
		return nil, common.ErrCannotListEntity(filter.TableName(), err)
	}
	return user, nil
}
