package restore

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) GetAppProfile(ctx context.Context) (*remodel.AppProfile, error) {
	result := remodel.AppProfile{}
	if err := s.db.Model(remodel.Amenity{}).Find(&result.Amenities).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := s.db.Model(remodel.RealEstatePostType{}).Find(&result.RealEstatePostTypes).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := s.db.Model(remodel.RealEstateType{}).Find(&result.RealEstateTypes).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
