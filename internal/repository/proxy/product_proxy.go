package proxy

import (
	"Erajaya/internal"
	"Erajaya/internal/entity"
	"Erajaya/internal/repository"
	"Erajaya/pkg/cache"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const (
	ProductCacheKey = "product_%s"
)

type ProductRepositoryProxy struct {
	repository.ProductRepository
	cache cache.RedisClientItf
}

func NewProductRepository(cache cache.RedisClientItf, productRepository repository.ProductRepository) internal.ProductRepositoryItf {
	return &ProductRepositoryProxy{
		ProductRepository: productRepository,
		cache:             cache,
	}
}

func (p *ProductRepositoryProxy) Store(ctx context.Context, product entity.Product) error {
	err := p.ProductRepository.Store(ctx, product)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryProxy) GetListSort(ctx context.Context, product entity.ListProductRequest) ([]entity.Product, error) {
	cacheKey := fmt.Sprintf(ProductCacheKey, product.SortBy)

	exist, _ := p.cache.Exists(ctx, cacheKey)

	if exist {
		var result []entity.Product
		res, err := p.cache.Get(ctx, cacheKey)

		err = json.Unmarshal(res, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	result, err := p.ProductRepository.GetListSort(ctx, product)
	if err != nil {
		return []entity.Product{}, err
	}

	res, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return result, nil
	}

	p.cache.SetEX(ctx, cacheKey, 300, string(res))

	return result, nil
}
