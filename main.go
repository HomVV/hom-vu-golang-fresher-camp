package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CREATE TABLE `Customers` (
// 	`Id` int(11) NOT NULL AUTO_INCREMENT,
// 	`Name` varchar(50) NOT NULL,
// 	`Address` varchar(100) NOT NULL,
// 	`PhoneNumber` varchar(12) NOT NULL,
// 	`Avatar` json DEFAULT NULL,
// 	`Email` varchar(100) DEFAULT NULL,
// 	`Status` tinyint(3) DEFAULT '1',
// 	`Created_At` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// 	`Updated_At` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`Id`),
// 	KEY `Email` (`Email`) USING BTREE
//   ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8
type Customer struct {
	Id          int    `json:"id,omitempty" gorm:"column:id;"`
	Name        string `json:"Name" gorm:"column:name;"`
	Address     string `json:"Address" gorm:"column:address;"`
	PhoneNumber string `json:"PhoneNumber" gorm:"column:phoneNumber;"`
	Email       string `json:"Email" gorm:"column:email;"`
}

func (Customer) TableName() string {
	return "Customers"
}
func main() {
	dsn := "food-delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	//insert
	newCustomer := Customer{
		Name:        "Ha",
		Address:     "Nghe An",
		PhoneNumber: "0382323289",
		Email:       "ha@gmail.com",
	}
	//db.Create(&newCustomer)
	if err := db.Create(&newCustomer).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(newCustomer)
}
