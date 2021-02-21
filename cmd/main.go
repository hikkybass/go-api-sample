package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yuki9541134/go-api-sample/infra"
	"github.com/yuki9541134/go-api-sample/usecase"
	"log"
	"net/http"
	"os"

	"github.com/yuki9541134/go-api-sample/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	productRepository := infra.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

    http.HandleFunc("/products", productHandler.HandleGetProducts)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
