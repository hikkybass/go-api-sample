package usecase

import (
	"github.com/yuki9541134/go-api-sample/domain/model"
	"testing"
)

type ProductRepositoryMock struct {}

func (pr ProductRepositoryMock) GetAll() ([]*model.Product, error) {
	products := []*model.Product {
		{
			ID: 1,
			Title: "aaa",
		},
	}
	return products, nil
}

func TestGetAllProduct(t *testing.T) {
	pr := ProductRepositoryMock{}
	pu := NewProductUseCase(pr)
	result, err := pu.GetAllProducts()

	if err != nil {
		t.Fatal(err)
	}

	if result[0].ID != 1 {
		t.Fatal(result)
	}
}
