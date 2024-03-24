package services

import (
	"Desktop/GoMark/app/models"
	"Desktop/GoMark/app/repository"
)

func GetAllProducts() []models.Product {
	products := repository.GetAllProducts()
	return products
}
