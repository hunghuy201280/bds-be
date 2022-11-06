package restore

import (
	"bds-service/common"
	"bds-service/common/entitycommon"
	"bds-service/modules/realestate/remodel"
	"context"
)

func (s *sqlStore) CreatePost(ctx context.Context, post *remodel.RealEstatePostCreate) (int, error) {
	transactionResult := s.db.Create(post)
	if err := transactionResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return post.Id, nil

}
