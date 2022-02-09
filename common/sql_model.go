package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Status    int        `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"Created_At"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"Updated_At"`
}
