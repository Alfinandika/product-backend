package internal

import (
	"Erajaya/internal/entity"
	"context"
)

type ProductServiceItf interface {
	AddProduct(ctx context.Context, req *entity.Product) (entity.Product, error)
	ListProduct(ctx context.Context, req *entity.ListProductRequest) ([]entity.Product, error)
}
