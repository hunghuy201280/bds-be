package userstorage

import (
	"bds-service/common"
	"bds-service/common/entitycommon"
	"bds-service/modules/user/usermodel"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) (int, error) {
	createResult := s.db.Create(&data)
	if err := createResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return data.Id, nil

}
