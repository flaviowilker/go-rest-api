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
	if products, err := p.ProductService.FindAll(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
	}
}

//FindByID ...
func (p *ProductAPI) FindByID(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		log.Println(errParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": errParam.Error()})
		return
	}

	if product, errFindByID := p.ProductService.FindByID(uint(id)); errFindByID != nil {
		log.Println(errFindByID)
		c.JSON(http.StatusNotFound, gin.H{"error": errFindByID.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
	}
}

//Create ...
func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	errBindJSON := c.BindJSON(&productDTO)
	if errBindJSON != nil {
		log.Println(errBindJSON)
		c.JSON(http.StatusBadRequest, gin.H{"error": errBindJSON.Error()})
		return
	}

	if product, errSave := p.ProductService.Create(ToProduct(productDTO)); errSave != nil {
		log.Println(errSave)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errSave.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"product": ToProductDTO(product)})
	}
}

//Update ...
func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO

	if err := c.BindJSON(&productDTO); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		log.Println(errParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": errParam.Error()})
		return
	}

	productDTO.ID = uint(id)

	if product, err := p.ProductService.Update(ToProduct(productDTO)); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
	}
}

//Delete ...
func (p *ProductAPI) Delete(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		log.Println(errParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": errParam.Error()})
		return
	}

	if product, err := p.ProductService.Delete(uint(id)); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
	}
}
