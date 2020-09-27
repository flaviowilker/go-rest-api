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
func (p *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	err := p.DB.Find(&products).Error

	return products, err
}

//FindByID ...
func (p *ProductRepository) FindByID(id uint) (Product, error) {
	var product Product
	err := p.DB.First(&product, id).Error

	return product, err
}

//Create ...
func (p *ProductRepository) Create(product Product) (Product, error) {
	err := p.DB.Create(&product).Error

	return product, err
}

//Update ...
func (p *ProductRepository) Update(product Product) (Product, error) {
	err := p.DB.Save(&product).Error

	return product, err
}

//Delete ...
func (p *ProductRepository) Delete(product Product) (Product, error) {
	err := p.DB.Delete(&product).Error

	return product, err
}
