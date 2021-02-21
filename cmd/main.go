package main

import (
	"github.com/joho/godotenv"
	"github.com/yuki9541134/go-api-sample/handler"
	"github.com/yuki9541134/go-api-sample/infra"
	"github.com/yuki9541134/go-api-sample/usecase"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB接続
	db := infra.ConnectDB()

	// 依存性の注入
	productRepository := infra.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

    http.HandleFunc("/products", productHandler.HandleGetProducts)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
