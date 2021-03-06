package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"toDoListGo/models"
)

var path = template.Must(template.ParseGlob("src/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	path.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	path.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("erro to convert price", err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("erro to convert quantity", err)
		}

		models.NewProduct(name, description, price, quantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.Delete(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetById(productId)
	path.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("erro to convert id", err)
		}
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("erro to convert price", err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("erro to convert quantity", err)
		}

		models.UpdateProduct(id, name, description, price, quantity)
	}
	http.Redirect(w, r, "/", 301)
}
