package product

import "github.com/jinzhu/gorm"

//Product ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
