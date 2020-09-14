package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	fmt.Println("start")
	connStr := "user=admin dbname=dev password=admin host=teste-postgres-compose sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
