package restore

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) ListRealEstateImages(
	ctx context.Context,
	reId int,
) ([]common.Image, error) {
	var result []common.Image
	db := s.db

	db = db.Table(remodel.RealEstateImage{}.TableName()).
		Where(
			common.JS{
				"re_id": reId,
			}).
		Joins("images on images.id = real_estate_images.id")

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil

}
