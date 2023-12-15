package usecase

import (
	"context"
	"github.com/meifanora/web-scaper-go/entity"
)

type ProductUseCase interface {
	InsertProducts(ctx context.Context, products []entity.Product) error
}