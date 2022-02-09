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

func CreateCustomer(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data customermodel.CustomerCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := customerstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := customerbiz.NewCreateCustomerBiz(store)

		if err := biz.CreateCustomer(c.Request.Context(), &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
