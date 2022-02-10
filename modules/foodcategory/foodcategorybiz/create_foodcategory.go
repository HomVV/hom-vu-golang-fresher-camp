package foodcategorybiz

import (
	"context"
	"demo/modules/foodcategory/foodcategorymodel"
)

type CreateFoodCategoryStore interface {
	CreateData(ctx context.Context, data *foodcategorymodel.CreateFoodCategory) error
}
type createFoodCategoryBiz struct {
	store CreateFoodCategoryStore
}

func NewCreateFoodCategoryBiz(store CreateFoodCategoryStore) *createFoodCategoryBiz {
	return &createFoodCategoryBiz{store: store}
}
func (biz *createFoodCategoryBiz) CreateFoodCategory(ctx context.Context,
	data *foodcategorymodel.CreateFoodCategory,
) error {
	if err := biz.store.CreateData(ctx, data); err != nil {
		return err
	}
	return nil
}
