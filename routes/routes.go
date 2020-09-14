package routes

import (
	"net/http"
	"toDoListGo/controllers"
)

func LoadRoat() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
