package routes

import (
	"net/http"
	"toDoListGo/controllers"
)

func LoadRoat() {
	http.HandleFunc("/", controllers.Index)
}
