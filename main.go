package main

import (
	"demo/component"
	"demo/modules/customer/customertransport/gincustomer"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "food-delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	// crud
	customers := r.Group("/customers")
	{
		customers.POST("", gincustomer.CreateCustomer(appCtx))

		// 	customers.GET("/:id", func(c *gin.Context) {
		// 		id, err := strconv.Atoi(c.Param("id"))
		// 		if err != nil {
		// 			c.JSON(401, map[string]interface{}{
		// 				"error": err.Error(),
		// 			})
		// 			return
		// 		}
		customers.GET("/:id", gincustomer.FindCustomer(appCtx))
		// 		var data Customer

		// 		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		// 			c.JSON(401, map[string]interface{}{
		// 				"error": err.Error(),
		// 			})
		// 			return
		// 		}
		// 		c.JSON(http.StatusOK, data)
		// 	})

		customers.GET("", gincustomer.ListCustomer(appCtx))
		customers.PATCH("/:id", gincustomer.UpdateCustomer(appCtx))
		customers.DELETE("/:id", gincustomer.SoftDeleteCustomer(appCtx))

		// 	customers.PATCH("/:id", func(c *gin.Context) {
		// 		id, err := strconv.Atoi(c.Param("id"))
		// 		if err != nil {
		// 			c.JSON(401, map[string]interface{}{
		// 				"error": err.Error(),
		// 			})
		// 			return
		// 		}

		// 		var data Customer
		// 		if err := c.ShouldBind(&data); err != nil {
		// 			c.JSON(401, map[string]interface{}{
		// 				"error": err.Error(),
		// 			})
		// 			return
		// 		}
		// 		if err := db.Where("id=?", id).Updates(&data).Error; err != nil {
		// 			c.JSON(401, map[string]interface{}{
		// 				"error": err.Error(),
		// 			})
		// 			return
		// 		}
		// 		c.JSON(http.StatusOK, data)
		// 	})
	}

	return r.Run()
}

// type Customer struct {
// 	Id          int    `json:"id,omitempty" gorm:"column:id;"`
// 	Name        string `json:"name" gorm:"column:name;"`
// 	Address     string `json:"address" gorm:"column:address;"`
// 	PhoneNumber string `json:"phonenumber" gorm:"column:phonenumber;"`
// 	Email       string `json:"email" gorm:"column:email;"`
// }

// func (Customer) TableName() string {
// 	return "Customers"
// }
