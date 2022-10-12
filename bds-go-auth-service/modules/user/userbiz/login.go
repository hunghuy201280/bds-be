package userbiz

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/component/tokenprovider"
	"bds-go-auth-service/modules/user/usermodel"
	"context"
	"errors"
	"time"
)

const (
	TokenExpiry        = int(15 * 24 * time.Hour)
	RefreshTokenExpiry = 2 * TokenExpiry
)

type LoginStore interface {
	FindOneUser(ctx context.Context, data *usermodel.UserFind) (*usermodel.UserLite, error)
}

type loginBiz struct {
	loginStore      LoginStore
	hasher          common.Hasher
	provider        tokenprovider.Provider
	refreshProvider tokenprovider.Provider
}

func NewLoginBiz(store LoginStore, hasher common.Hasher, provider tokenprovider.Provider, refreshProvider tokenprovider.Provider) *loginBiz {
	return &loginBiz{
		store,
		hasher,
		provider,
		refreshProvider,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, *tokenprovider.Token, error) {
	err := data.Validate()
	if err != nil {
		return nil, nil, common.ErrLoginFailed(err)
	}
	store := biz.loginStore
	user, err := store.FindOneUser(ctx, &usermodel.UserFind{
		Phone: data.Phone,
	})
	if err != nil {
		return nil, nil, common.ErrLoginFailed(errors.New("wrong username or password"))
	}

	hasher := biz.hasher

	if hasher.ValidatePassword(user.Password, data.Password+user.Salt) {
		return nil, nil, common.ErrLoginFailed(errors.New("wrong username or password"))
	}
	token, err := biz.provider.Generate(tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.String(),
	}, TokenExpiry)

	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := biz.refreshProvider.Generate(tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role.String(),
	}, RefreshTokenExpiry)

	return token, refreshToken, nil

}
