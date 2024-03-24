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
