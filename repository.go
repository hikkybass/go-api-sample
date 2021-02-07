package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getProducts() []Product {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	query := `select id, title from products limit 10;`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	var products []Product
	for rows.Next() {
		var (product Product)
		err := rows.Scan(&product.ID, &product.Title)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products
}
