package main

import (
	"fmt"
	"html/template"
	"net/http"
	"toDoListGo/models"
)

var path = template.Must(template.ParseGlob("src/*.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	path.ExecuteTemplate(w, "Index", products)
}
