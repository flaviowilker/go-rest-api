package product

import "github.com/jinzhu/gorm"

//Product ...
type Product struct {
	gorm.Model
	Code        string  `gorm:"column:code;type:varchar(50)"`
	Price       float64 `gorm:"column:price;"`
	Description string  `gorm:"column:description;type:varchar(100)"`
}

//TableName ...
func (Product) TableName() string {
	return "t_product"
}
