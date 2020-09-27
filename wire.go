//+build wireinject

package main

import (
	"github.com/flaviowilker/go-rest-api/product"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
