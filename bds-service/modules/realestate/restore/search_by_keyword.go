package restore

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) SearchByKeyword(
	ctx context.Context,
	keyword string,
) ([]remodel.RealEstate, error) {
	var result []remodel.RealEstate
	db := s.db

	db = db.Table(remodel.RealEstate{}.TableName())
	db = db.Raw("SELECT r.*,  " +
		"   p.name_en AS p_name, " +
		"  d.name_en AS d_name, " +
		"  w.name_en AS w_name FROM real_estates r " +
		"JOIN provinces AS p ON p.code = r.province_id  " +
		"JOIN districts AS d ON d.code = r.district_id  " +
		"JOIN wards AS w ON w.code = r.ward_id " +
		"WHERE p.name_en LIKE '%" + keyword + "%' " +
		"OR d.name_en LIKE '%" + keyword + "%' " +
		"OR w.name_en LIKE '%" + keyword + "%' " +
		"OR r.address LIKE '%" + keyword + "%' ",
	)

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil

}
