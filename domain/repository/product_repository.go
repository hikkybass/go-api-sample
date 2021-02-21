package repository

import (
	"github.com/yuki9541134/go-api-sample/domain/model"
)

type ProductRepository interface {
	GetAll() ([]*model.Product, error)
}
