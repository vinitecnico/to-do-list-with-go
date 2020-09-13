package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectDB() *sql.DB {
	fmt.Println("start")
	connStr := "user=admin dbname=dev password=admin host=teste-postgres-compose sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	productID   int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var path = template.Must(template.ParseGlob("src/*.html"))

func main() {
	db := conectDB()
	defer db.Close()
	http.HandleFunc("/", index)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDB()
	rows, err := db.Query("SELECT * from products")
	if err != nil {
		panic(err.Error())
	}
	products := []Product{}

	for rows.Next() {
		var productID, quantity int
		var name, description string
		var price float64
		err = rows.Scan(&productID, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		product := Product{productID, name, description, price, quantity}
		products = append(products, product)
	}

	path.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
