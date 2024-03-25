package controllers

import (
	"Desktop/GoMark/app/repository"
	"Desktop/GoMark/app/services"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := services.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert price error: ", err)
		}
		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Convert quantity error: ", err)
		}

		services.CreateProduct(name, description, convertPrice, convertQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := services.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Convert id error: ", err)
		}

		convertPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert price error: ", err)
		}

		convertQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Convert quantity error: ", err)
		}

		repository.UpdateProduct(convertId, name, description, convertPrice, convertQuantity)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
