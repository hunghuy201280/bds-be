package rebiz

import (
	"bds-service/common"
	"bds-service/modules/realestate/remodel"
	"context"
	"errors"
	"github.com/gookit/goutil"
)

type SearchByKeywordStorage interface {
	SearchByKeyword(ctx context.Context, keyword string) ([]remodel.RealEstate, error)
}

type searchByKeywordBiz struct {
	store SearchByKeywordStorage
}

func NewSearchByKeywordBiz(storage SearchByKeywordStorage) *searchByKeywordBiz {
	return &searchByKeywordBiz{
		store: storage,
	}
}

func (biz *searchByKeywordBiz) SearchByKeyword(ctx context.Context, keyword string) ([]remodel.RealEstate, error) {
	if goutil.IsEmpty(keyword) {
		return nil, common.ErrInvalidRequest(errors.New("data is empty"))
	}

	store := biz.store
	result, err := store.SearchByKeyword(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return result, nil

}
