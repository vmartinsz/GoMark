package repository

import (
	"Desktop/GoMark/app/db"
	models "Desktop/GoMark/app/models"
)

func GetAllProducts() []models.Product {
	db := db.Connect()

	result, err := db.Query("SELECT * FROM product order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := models.Product{}
	products := []models.Product{}

	for result.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.Connect()

	result, err := db.Prepare("INSERT INTO product (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	result.Exec(name, description, price, quantity)
	defer db.Close()
}

func EditProduct(id string) models.Product {
	db := db.Connect()

	productBank, err := db.Query("SELECT * FROM product WHERE id=$1", id)
	if err != nil {
		panic("Error EditProduct repository" + err.Error())
	}

	readyToUpdate := models.Product{}

	for productBank.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productBank.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		readyToUpdate.Id = id
		readyToUpdate.Name = name
		readyToUpdate.Description = description
		readyToUpdate.Price = price
		readyToUpdate.Quantity = quantity
	}
	defer db.Close()
	return readyToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.Connect()

	atualizaProduto, err := db.Prepare("UPDATE product set name=$1, description=$2, price=$3, quantity=$4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(name, description, price, quantity, id)
	defer db.Close()
}
