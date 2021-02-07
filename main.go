package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    http.HandleFunc("/products", handleGetProducts)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
