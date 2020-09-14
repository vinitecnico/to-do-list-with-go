package models

import "toDoListGo/db"

type Product struct {
	productID   int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAll() []Product {
	db := db.ConectDB()
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

	defer db.Close()
	return products
}

func NewProduct(name string, description string, price float64, quantity int) {
	db := db.ConectDB()
	insertdata, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertdata.Exec(name, description, price, quantity)
	defer db.Close()
}
