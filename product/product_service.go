package product

import (
	"errors"
	"log"
)

//ProductService ...
type ProductService struct {
	ProductRepository ProductRepository
}

//ProvideProductService ...
func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

//FindAll ...
func (p *ProductService) FindAll() ([]Product, error) {
	returnValue, err := p.ProductRepository.FindAll()

	return returnValue, err
}

//FindByID ...
func (p *ProductService) FindByID(id uint) (Product, error) {
	returnValue, err := p.ProductRepository.FindByID(id)

	return returnValue, err
}

//Create ...
func (p *ProductService) Create(product Product) (Product, error) {
	returnValue, err := p.ProductRepository.Create(product)

	return returnValue, err
}

//Update ...
func (p *ProductService) Update(product Product) (Product, error) {
	if productFinded, err := p.ProductRepository.FindByID(uint(product.ID)); err != nil {
		log.Println(err)
		return Product{}, err
	} else if productFinded == (Product{}) {
		log.Println("Update: Not Found")
		return Product{}, errors.New("Update: Not Found")
	}

	returnValue, errSave := p.ProductRepository.Update(product)

	return returnValue, errSave
}

//Delete ...
func (p *ProductService) Delete(id uint) (Product, error) {
	product, errFindByID := p.ProductRepository.FindByID(id)
	if errFindByID != nil {
		log.Println(errFindByID)
		return Product{}, errFindByID
	}

	if product == (Product{}) {
		log.Println("Delete: Not Found")
		return Product{}, errors.New("Delete: Not Found")
	}

	returnValue, err := p.ProductRepository.Delete(product)

	return returnValue, err
}
