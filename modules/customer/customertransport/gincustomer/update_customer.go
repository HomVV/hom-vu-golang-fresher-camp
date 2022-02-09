package gincustomer

import (
	"demo/common"
	"demo/component"
	"demo/modules/customer/customerbiz"
	"demo/modules/customer/customermodel"
	"demo/modules/customer/customerstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		var data customermodel.CustomerUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := customerstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := customerbiz.NewUpdateCustomerBiz(store)

		if err := biz.UpdateCustomer(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
