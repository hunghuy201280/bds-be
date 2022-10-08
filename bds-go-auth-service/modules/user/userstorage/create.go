package userstorage

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/common/entitycommon"
	"bds-go-auth-service/modules/user/usermodel"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) (int, error) {
	createResult := s.db.Create(&data)
	if err := createResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return data.Id, nil

}
