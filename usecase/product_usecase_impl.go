package usecase

import (
	"context"
	"database/sql"
	"github.com/meifanora/web-scaper-go/entity"
	"github.com/meifanora/web-scaper-go/repository"
)

type ProductUseCaseImpl struct {
	productRepository repository.ProductRepository
	database          *sql.DB
}

func NewProductUseCase(productRepository repository.ProductRepository, database *sql.DB) ProductUseCase {
	return &ProductUseCaseImpl{
		productRepository: productRepository,
		database:          database,
	}
}

func (c *ProductUseCaseImpl) InsertProducts(ctx context.Context, products []entity.Product) error {
	tx, err := c.database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, product := range products {
		err = c.productRepository.Insert(ctx, tx, product)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
