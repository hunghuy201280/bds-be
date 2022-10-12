package userbiz

import (
	"bds-go-auth-service/component/tokenprovider"
	"bds-go-auth-service/modules/user/usermodel"
	"context"
)

type refreshTokenBiz struct {
	refreshProvider tokenprovider.Provider
	provider        tokenprovider.Provider
}

func NewRefreshTokenBiz(refreshProvider tokenprovider.Provider, provider tokenprovider.Provider) *refreshTokenBiz {
	return &refreshTokenBiz{
		provider:        provider,
		refreshProvider: refreshProvider,
	}
}

func (biz *refreshTokenBiz) RefreshToken(ctx context.Context, data *usermodel.UserRefreshToken) (token *tokenprovider.Token, newRefreshToken *tokenprovider.Token, err error) {
	err = data.Validate()
	if err != nil {
		return nil, nil, err
	}
	tokenPayload, err := biz.refreshProvider.Validate(data.RefreshToken)
	if err != nil || tokenPayload == nil {
		return nil, nil, err
	}
	if tokenPayload.UserId != data.UserId {
		return nil, nil, tokenprovider.ErrInvalidToken
	}

	token, err = biz.provider.Generate(*tokenPayload, TokenExpiry)
	if err != nil {
		return nil, nil, err
	}

	newRefreshToken, err = biz.refreshProvider.Generate(*tokenPayload, RefreshTokenExpiry)
	if err != nil {
		return nil, nil, err
	}

	return token, newRefreshToken, nil

}
