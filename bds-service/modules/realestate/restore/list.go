package restore

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
	"github.com/gookit/goutil"
	"strconv"
)

func (s *sqlStore) ListRealEstate(
	ctx context.Context,
	conditions map[string]any,
	filter *remodel.Filter,
	paging *common.Paging,
	moreKeys ...string,

) ([]remodel.RealEstate, error) {
	var result []remodel.RealEstate
	db := s.db

	for _, key := range moreKeys {
		db = db.Preload(key)
	}
	db = db.Table(remodel.RealEstate{}.TableName()).Where(conditions)

	if v := filter.RealEstateAmenityTypeIds; !goutil.IsEmpty(v) {
		db.Preload("Amenities", "amenityId in ?", v)
	} else {
		db.Preload("Amenities")
	}
	db.Preload("RealEstateType")
	db.Preload("Images")
	if filter != nil {
		if v := filter.MinPrice; v != 0 {
			db.Where("min_price >= ?", v)
		}
		if v := filter.MaxPrice; v != 0 {
			db.Where("`max_price` <= ?", v)
		}
		if v := filter.NoBedrooms; v != 0 {
			db.Where("`no_bedrooms` = ?", v)
		}
		if v := filter.NoWC; v != 0 {
			db.Where("`no_wc` = ?", v)
		}
		if v := filter.MinArea; v != 0 {
			db.Where("`area` >= ?", v)
		}
		if v := filter.MaxArea; v != 0 {
			db.Where("`area` <= ?", v)
		}

		if v := filter.Floors; v != 0 {
			db.Where("`floors` = ?", v)
		}
		if v := filter.RealEstateTypeIds; !goutil.IsEmpty(v) {
			db.Where("`real_estate_type_id` in ? ", v)
		}

		if v := filter.ProvinceId; !goutil.IsEmpty(v) {
			db.Where("`province_id` = ? ", v)
		}
		if v := filter.DistrictId; !goutil.IsEmpty(v) {
			db.Where("`district_id` = ? ", v)
		}
		if v := filter.WardId; !goutil.IsEmpty(v) {
			db.Where("`ward_id` = ? ", v)
		}
	}

	if paging == nil {
		paging = &common.Paging{}
	}

	paging.Fulfill()

	if v := paging.FakeCursor; v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			db = db.Where("id < ?", id)
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	paging.Total = int64(len(result))

	return result, nil

}
