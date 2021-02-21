package infra

import (
	"database/sql"
	"github.com/yuki9541134/go-api-sample/domain/model"
	"github.com/yuki9541134/go-api-sample/domain/repository"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
	return productRepository{db}
}

func (pr productRepository) GetAll() ([]*model.Product, error) {
	query := `select id, title from products limit 10;`
	rows, err := pr.db.Query(query)
	if err != nil {
		panic(err)
	}

	var products []*model.Product
	for rows.Next() {
		var (product model.Product)
		err := rows.Scan(&product.ID, &product.Title)
		if err != nil {
			panic(err)
		}
		products = append(products, &product)
	}

	return products, nil
}
