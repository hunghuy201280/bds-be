package userbiz

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/common/entitycommon"
	"bds-go-auth-service/common/l"
	"bds-go-auth-service/modules/user/usermodel"
	"context"
	"errors"
)

var ll = l.New()

type RegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (int, error)
	FindOneUser(ctx context.Context, data *usermodel.UserFind) (*usermodel.UserLite, error)
}

type registerBiz struct {
	registerStore RegisterStore
	hasher        common.Hasher
}

func NewRegisterBiz(store RegisterStore, hasher common.Hasher) *registerBiz {
	return &registerBiz{
		store,
		hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserLite, error) {
	err := data.Validate()
	if err != nil {
		return nil, common.ErrCannotCreateEntity(data.TableName(), err)
	}
	store := biz.registerStore
	user, err := store.FindOneUser(ctx, &usermodel.UserFind{
		Phone: data.Phone,
	})
	if user != nil {
		return nil, common.ErrEntityExisted(data.TableName(), errors.New("entity existed"))
	}
	hasher := biz.hasher
	salt, err := hasher.HashPassword(common.RandStringBytes(10))
	if err != nil {
		return nil, common.ErrCannotCreateEntity(data.TableName(), err)
	}
	rawPassword := data.Password
	data.Password, err = hasher.HashPassword(rawPassword + salt)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(data.TableName(), err)
	}

	data.Role = usermodel.BUYER.String()
	data.Status = entitycommon.NORMAL
	data.Salt = salt
	userId, err := store.CreateUser(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(data.TableName(), err)
	}
	return &usermodel.UserLite{
		SQLModel: common.SQLModel{
			Id: userId,
		},
		Phone:     data.Phone,
		LastName:  data.LastName,
		FirstName: data.FirstName,
		Role:      usermodel.UserRole(data.Role),
	}, nil

}
