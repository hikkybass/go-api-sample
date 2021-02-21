package handler

import (
	"fmt"
	"github.com/yuki9541134/go-api-sample/domain/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type ProductUseCaseMock struct {}

func (pu ProductUseCaseMock) GetAllProducts() ([]*model.Product, error) {
	products := []*model.Product {
		{
			ID: 1,
			Title: "aaa",
		},
	}

	return products, nil
}

func TestHandleGetProducts(t *testing.T) {
	pu := ProductUseCaseMock{}
	ph := NewProductHandler(pu)

	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	ph.HandleGetProducts(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}
	fmt.Println(resp.Header.Get("Content-Type"))

	fmt.Println(ioutil.ReadAll(resp.Body))
}
