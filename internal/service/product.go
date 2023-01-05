package service

import (
	"Erajaya/internal"
	"Erajaya/internal/entity"
	"context"
)

type ProductService struct {
	productRepo internal.ProductRepositoryItf
}

func NewProductService(productRepo internal.ProductRepositoryItf) internal.ProductServiceItf {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (p *ProductService) AddProduct(ctx context.Context, req *entity.Product) (entity.Product, error) {
	err := p.productRepo.Store(ctx, *req)
	if err != nil {
		return entity.Product{}, err
	}

	return *req, nil
}

func (p *ProductService) ListProduct(ctx context.Context, req *entity.ListProductRequest) ([]entity.Product, error) {
	products, err := p.productRepo.GetListSort(ctx, *req)
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}
