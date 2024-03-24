package controllers

import (
	"Desktop/GoMark/app/services"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := services.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
