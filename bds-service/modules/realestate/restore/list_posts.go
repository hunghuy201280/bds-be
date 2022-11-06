package restore

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
	"github.com/gookit/goutil"
	"strconv"
)

func (s *sqlStore) ListRealEstatePost(
	ctx context.Context,
	conditions map[string]any,
	filter *remodel.Filter,
	paging *common.Paging,
	moreKeys ...string,

) ([]remodel.RealEstatePost, error) {
	var result []remodel.RealEstatePost
	db := s.db

	for _, key := range moreKeys {
		db = db.Preload(key)
	}
	db = db.Table(remodel.RealEstatePost{}.TableName()).Where(conditions)
	db.Preload("RePostType")
	if filter != nil {
		if v := filter.MinPrice; v != 0 {
			db.Preload("RealEstate", "price >= ?", v)
		}
		if v := filter.MaxPrice; v != 0 {
			db.Preload("RealEstate", "`price` <= ?", v)
		}
		if v := filter.NoBedrooms; v != 0 {
			db.Preload("RealEstate", "`no_bedrooms` = ?", v)
		}
		if v := filter.NoWC; v != 0 {
			db.Preload("RealEstate", "`no_wc` = ?", v)
		}
		if v := filter.MinArea; v != 0 {
			db.Preload("RealEstate", "`area` >= ?", v)
		}
		if v := filter.MaxArea; v != 0 {
			db.Preload("RealEstate", "`area` <= ?", v)
		}

		if v := filter.Floors; v != 0 {
			db.Preload("RealEstate", "`floors` = ?", v)
		}
		if v := filter.RealEstateTypeIds; !goutil.IsEmpty(v) {
			db.Preload("RealEstate", "`re_type_id` in ? ", v)
		}

		if v := filter.ProvinceId; !goutil.IsEmpty(v) {
			db.Preload("RealEstate", "`province_id` = ? ", v)
		}
		if v := filter.DistrictId; !goutil.IsEmpty(v) {
			db.Preload("RealEstate", "`district_id` = ? ", v)
		}
		if v := filter.WardId; !goutil.IsEmpty(v) {
			db.Preload("RealEstate", "`ward_id` = ? ", v)
		}
	}

	if v := filter.RealEstateAmenityTypeIds; !goutil.IsEmpty(v) {
		db.Preload("RealEstate.Amenities", "Id in ?", v)
	} else {
		db.Preload("RealEstate.Amenities")
	}
	db.Preload("RealEstate.RealEstateType")
	db.Preload("RealEstate.Images")

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
