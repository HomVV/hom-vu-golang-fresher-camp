package customermodel

type Customer struct {
	Id          int    `json:"id,omitempty" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Address     string `json:"address" gorm:"column:address;"`
	PhoneNumber string `json:"phonenumber" gorm:"column:phonenumber;"`
	Email       string `json:"email" gorm:"column:email;"`
}

func (Customer) TableName() string {
	return "Customers"
}

type CustomerUpdate struct {
	Name        *string `json:"name" gorm:"column:name;"`
	Address     *string `json:"address" gorm:"column:address;"`
	PhoneNumber *string `json:"phonenumber" gorm:"column:phonenumber;"`
	Email       *string `json:"email" gorm:"column:email;"`
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
}

func (CustomerCreate) TableName() string {
	return Customer{}.TableName()
}
