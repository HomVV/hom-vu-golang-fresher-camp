package gincustomer

import (
	"demo/common"
	"demo/component"
	"demo/modules/customer/customerbiz"
	"demo/modules/customer/customerstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SoftDeleteCustomer(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
		}
		store := customerstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := customerbiz.NewSoftDeleteCustomerBiz(store)

		if err := biz.SoftDeleteCustomer(c.Request.Context(), id); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
