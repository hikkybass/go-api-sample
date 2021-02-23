package handler

import (
	"encoding/json"
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
		t.Fatalf(`resp.StatusCode = %d, want 200`, resp.StatusCode)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var products []model.Product
	if err := json.Unmarshal(respBodyBytes, &products); err != nil {
		panic(err)
	}
	if products[0].ID != 1 {
		t.Fatalf(`products[0].ID = %d, want 1`, products[0].ID)
	}
}
