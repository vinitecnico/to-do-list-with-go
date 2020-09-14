package controllers

import (
	"html/template"
	"net/http"
	"toDoListGo/models"
)

var path = template.Must(template.ParseGlob("src/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	path.ExecuteTemplate(w, "Index", products)
}
