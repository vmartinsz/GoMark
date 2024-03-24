package routes

import (
	"Desktop/GoMark/app/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
