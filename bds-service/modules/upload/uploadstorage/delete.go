package uploadstorage

import (
	"bds-service/common"
	"context"
)

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db
	if err := db.
		Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
