package foodcategorybiz

import (
	"context"
	"demo/common"
	"demo/modules/foodcategory/foodcategorymodel"
)

type GetFoodCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*foodcategorymodel.FoodCategory, error)
}

type getFoodCategoryBiz struct {
	store GetFoodCategoryStore
}

func NewFoodCategoryBiz(store GetFoodCategoryStore) *getFoodCategoryBiz {
	return &getFoodCategoryBiz{store: store}
}
func (biz *getFoodCategoryBiz) GetFoodCategory(
	ctx context.Context,
	id int) (*foodcategorymodel.FoodCategory, error) {
	var data *foodcategorymodel.FoodCategory
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(foodcategorymodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(foodcategorymodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(foodcategorymodel.EntityName, nil)
	}
	return data, nil
}
