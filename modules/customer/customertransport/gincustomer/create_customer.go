package gincustomer

import (
	"demo/modules/customer/customerbiz"
	"demo/modules/customer/customermodel"
	"demo/modules/customer/customerstorage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data customermodel.CustomerCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := customerstorage.NewSQLStore(db)

		biz := customerbiz.NewCreateCustomerBiz(store)

		if err := biz.CreateCustomer(c.Request.Context(), &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
