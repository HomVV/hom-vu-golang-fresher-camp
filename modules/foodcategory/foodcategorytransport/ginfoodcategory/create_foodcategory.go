package ginfoodcategory

import (
	"demo/common"
	"demo/component"
	"demo/modules/foodcategory/foodcategorybiz"
	"demo/modules/foodcategory/foodcategorymodel"
	"demo/modules/foodcategory/foodcategorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodcategorymodel.CreateFoodCategory
		if err := c.ShouldBind(&data); err != nil {
			common.ErrInvalidRequest(err)
		}
		store := foodcategorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodcategorybiz.NewCreateFoodCategoryBiz(store)

		if err := biz.CreateFoodCategory(c.Request.Context(), &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
