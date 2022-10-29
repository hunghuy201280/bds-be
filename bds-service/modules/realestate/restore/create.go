package restore

import (
	"bds-service/common"
	"bds-service/common/entitycommon"
	"bds-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) CreateRealEstate(ctx context.Context, data *remodel.RealEstateCreate) (int, error) {
	data.SQLModel = common.SQLModel{Status: entitycommon.NORMAL}
	createResult := s.db.Create(&data)
	if err := createResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return data.Id, nil
}
