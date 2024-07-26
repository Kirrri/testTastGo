package services

import "OneGO/models"

var products []models.Product
var currentProductID int

func GetProducts() []models.Product {
	return products
}

func GetProductByID(id int) (models.Product, bool) {
	for _, item := range products {
		if item.Id == id {
			return item, true
		}
	}
	return models.Product{}, false
}

func CreateProduct(product *models.Product) {
	currentProductID++
	product.Id = currentProductID
	products = append(products, *product)
}

func UpdateProduct(id int, product *models.Product) (models.Product, bool) {
	for index, item := range products {
		if item.Id == id {
			products[index] = *product
			product.Id = id
			return *product, true
		}
	}
	return models.Product{}, false
}

func DeleteProduct(id int) bool {
	for index, item := range products {
		if item.Id == id {
			products = append(products[:index], products[index+1:]...)
			return true
		}
	}
	return false
}
