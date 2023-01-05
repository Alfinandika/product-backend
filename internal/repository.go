package internal

import (
	"Erajaya/internal/entity"
	"context"
)

type ProductRepositoryItf interface {
	Store(ctx context.Context, product entity.Product) error
	GetListSort(ctx context.Context, product entity.ListProductRequest) ([]entity.Product, error)
}
