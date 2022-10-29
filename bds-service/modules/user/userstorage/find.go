package userstorage

import (
	"bds-service/common"
	"bds-service/modules/user/usermodel"
	"context"
)

func (s *sqlStore) FindOneUser(ctx context.Context, data *usermodel.UserFind) (*usermodel.UserLite, error) {
	db := s.db

	var user usermodel.UserLite
	if err := db.Where(&data).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil

}
