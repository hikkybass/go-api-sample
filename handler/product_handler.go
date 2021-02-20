package handler

import (
	"encoding/json"
	"github.com/go-api-sample/usecase"
	"io"
	"net/http"
)

type ProductHandler interface {
	HandleGetProducts(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) ProductHandler {
	return productHandler{useCase}
}

func (ph productHandler) HandleGetProducts (w http.ResponseWriter, r *http.Request) {
	products, err := ph.productUseCase.GetAllProducts()
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(res))
}
