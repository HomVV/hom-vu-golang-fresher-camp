package main

import (
	"demo/component"
	"demo/middleware"
	"demo/modules/customer/customertransport/gincustomer"
	"demo/modules/foodcategory/foodcategorytransport/ginfoodcategory"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//dsn := "food-delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// crud
	customers := r.Group("/customers")
	{
		customers.POST("", gincustomer.CreateCustomer(appCtx))
		customers.GET("/:id", gincustomer.FindCustomer(appCtx))
		customers.GET("", gincustomer.ListCustomer(appCtx))
		customers.PATCH("/:id", gincustomer.UpdateCustomer(appCtx))
		customers.DELETE("/:id", gincustomer.SoftDeleteCustomer(appCtx))

	}
	foodCategories := r.Group("/foodCategories")
	{
		foodCategories.POST("", ginfoodcategory.CreateFood(appCtx))
		foodCategories.GET(":id", ginfoodcategory.GetFoodCategory(appCtx))
	}

	return r.Run()
}
