package customermodel

import (
	"demo/common"
	"errors"
	"strings"
)

const EntityName = "customer"

type Customer struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:address;"`
	PhoneNumber     string `json:"phonenumber" gorm:"column:phonenumber;"`
	Email           string `json:"email" gorm:"column:email;"`
}

func (Customer) TableName() string {
	return "Customers"
}

type CustomerUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name;"`
	Address         *string `json:"address" gorm:"column:address;"`
	PhoneNumber     *string `json:"phonenumber" gorm:"column:phonenumber;"`
	Email           *string `json:"email" gorm:"column:email;"`
}

func (CustomerUpdate) TableName() string {
	return Customer{}.TableName()
}

type CustomerCreate struct {
	Id          int    `json:"id,omitempty" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Address     string `json:"address" gorm:"column:address;"`
	PhoneNumber string `json:"phonenumber" gorm:"column:phonenumber;"`
	Email       string `json:"email" gorm:"column:email;"`
	//Status      int8   `json:"status" gorm:"column:status"`
}

func (CustomerCreate) TableName() string {
	return Customer{}.TableName()
}

func (res *CustomerCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return errors.New("Customer name can't be blank")
	}
	return nil
}
