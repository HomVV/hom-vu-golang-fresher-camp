package foodcategorymodel

import (
	"demo/common"
)

const EntityName = "FoodCategory"

type FoodCategory struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Description     string `json:"description" gorm:"column:description;"`
}

func (FoodCategory) TableName() string {
	return "FoodCategories"
}

type UpdateFoodCategory struct {
	Name        *string `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
}

func (UpdateFoodCategory) TableName() string {
	return FoodCategory{}.TableName()
}

type CreateFoodCategory struct {
	Id          int    `json:"id,omitempty" gorm:"id"`
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (CreateFoodCategory) TableName() string {
	return FoodCategory{}.TableName()
}
