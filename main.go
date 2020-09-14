package main

import (
	"fmt"
	"net/http"
	"toDoListGo/routes"
)

func main() {
	routes.LoadRoat()
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
