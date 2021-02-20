package repository

import (
	"github.com/go-api-sample/domain/model"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
}
