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

func FindCustomer(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		store := customerstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := customerbiz.NewFindCustomerBiz(store)

		result, err := biz.FindCustomer(c.Request.Context(), id)
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
