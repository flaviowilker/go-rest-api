package product

//ProductService ...
type ProductService struct {
	ProductRepository ProductRepository
}

//ProvideProductService ...
func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

//FindAll ...
func (p *ProductService) FindAll() []Product {
	return p.ProductRepository.FindAll()
}

//FindByID ...
func (p *ProductService) FindByID(id uint) Product {
	return p.ProductRepository.FindByID(id)
}

//Save ...
func (p *ProductService) Save(product Product) Product {
	p.ProductRepository.Save(product)

	return product
}

//Delete ...
func (p *ProductService) Delete(product Product) {
	p.ProductRepository.Delete(product)
}
