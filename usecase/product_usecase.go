package usecase

import (
	"github.com/yuki9541134/go-api-sample/domain/model"
	"github.com/yuki9541134/go-api-sample/domain/repository"
)

type ProductUseCase interface {
	GetAllProducts() ([]model.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return productUseCase{pr}
}

func (pu productUseCase) GetAllProducts() ([]model.Product, error) {
	return pu.productRepository.GetAll()
}
