package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ProductAPI ...
type ProductAPI struct {
	ProductService ProductService
}

//ProvideProductAPI ...
func ProvideProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

//FindAll ...
func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

//FindByID ...
func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

//Create ...
func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductService.Save(ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createdProduct)})
}

//Update ...
func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

//Delete ...
func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}