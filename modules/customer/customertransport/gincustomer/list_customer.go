package gincustomer

import (
	"demo/common"
	"demo/component"
	"demo/modules/customer/customerbiz"
	"demo/modules/customer/customermodel"
	"demo/modules/customer/customerstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCustomer(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter customermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}
		paging.Fulfill()

		store := customerstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := customerbiz.NewListCustomerBiz(store)
		result, err := biz.ListCustomer(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
