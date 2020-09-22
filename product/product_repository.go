package product

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ProductRepository ...
type ProductRepository struct {
	DB *gorm.DB
}

//ProvideProductRepostiory ...
func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

//FindAll ...
func (p *ProductRepository) FindAll() []Product {
	var products []Product
	p.DB.Find(&products)

	return products
}

//FindByID ...
func (p *ProductRepository) FindByID(id uint) Product {
	var product Product
	p.DB.First(&product, id)

	return product
}

//Save ...
func (p *ProductRepository) Save(product Product) Product {
	p.DB.Save(&product)

	return product
}

//Delete ...
func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}
