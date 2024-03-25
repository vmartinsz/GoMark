package services

import (
	"Desktop/GoMark/app/models"
	"Desktop/GoMark/app/repository"
)

func GetAllProducts() []models.Product {
	products := repository.GetAllProducts()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	repository.CreateProduct(name, description, price, quantity)
}

func EditProduct(id string) models.Product {
	readyToUpdate := repository.EditProduct(id)
	return readyToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	repository.UpdateProduct(id, name, description, price, quantity)
}
