package uploadstorage

import (
	"bds-service/common"
	"context"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := store.db
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
