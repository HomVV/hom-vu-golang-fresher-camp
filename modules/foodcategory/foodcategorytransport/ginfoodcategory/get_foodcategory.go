package ginfoodcategory

import (
	"demo/common"
	"demo/component"
	"demo/modules/foodcategory/foodcategorybiz"
	"demo/modules/foodcategory/foodcategorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetFoodCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.ErrInvalidRequest(err)
		}
		store := foodcategorystorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := foodcategorybiz.NewFoodCategoryBiz(store)

		data, err := biz.GetFoodCategory(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
