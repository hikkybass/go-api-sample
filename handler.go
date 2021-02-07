package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handleGetProducts (w http.ResponseWriter, r *http.Request) {
	products := getProductsUseCase()

	res, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(res))
}
