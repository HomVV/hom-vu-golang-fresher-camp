package foodcategorystorage

import (
	"context"
	"demo/common"
	"demo/modules/foodcategory/foodcategorymodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*foodcategorymodel.FoodCategory, error) {
	var data foodcategorymodel.FoodCategory
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(foodcategorymodel.FoodCategory{}.TableName()).Where(condition)
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
