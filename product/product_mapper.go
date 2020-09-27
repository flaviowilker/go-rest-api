package product

import "github.com/jinzhu/gorm"

//ToProduct ...
func ToProduct(productDTO ProductDTO) Product {
	return Product{
		Model: gorm.Model{
			ID: productDTO.ID,
		},
		Code:        productDTO.Code,
		Price:       productDTO.Price,
		Description: productDTO.Description,
	}
}

//ToProductDTO ...
func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{
		ID:          product.ID,
		Code:        product.Code,
		Price:       product.Price,
		Description: product.Description,
	}
}

//ToProductDTOs ...
func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}

//ToProducts ...
func ToProducts(productDTOs []ProductDTO) []Product {
	products := make([]Product, len(productDTOs))

	for i, itm := range productDTOs {
		products[i] = ToProduct(itm)
	}

	return products
}
