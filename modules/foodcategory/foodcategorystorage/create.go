package foodcategorystorage

import (
	"context"
	"demo/common"
	"demo/modules/foodcategory/foodcategorymodel"
)

func (s *sqlStore) CreateData(ctx context.Context, data *foodcategorymodel.CreateFoodCategory) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
