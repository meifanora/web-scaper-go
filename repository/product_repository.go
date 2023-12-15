package repository

import (
	"context"
	"database/sql"
	"github.com/meifanora/web-scaper-go/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product entity.Product) error
}
