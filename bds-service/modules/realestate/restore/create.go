package restore

import (
	"bds-go-auth-service/common"
	"bds-go-auth-service/common/entitycommon"
	"bds-go-auth-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) CreateRealEstate(ctx context.Context, data *remodel.RealEstateCreate) (int, error) {
	createResult := s.db.Create(&data)
	if err := createResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return data.Id, nil
}
