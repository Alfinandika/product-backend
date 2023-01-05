package repository

import (
	"Erajaya/internal/entity"
	"Erajaya/pkg/database"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type ProductRepository struct {
	client database.PostgreClientItf
}

func NewProductRepository(client database.PostgreClientItf) *ProductRepository {
	return &ProductRepository{
		client: client,
	}
}

func (p *ProductRepository) Store(ctx context.Context, product entity.Product) error {
	db, err := p.client.GetMaster(ctx)
	if err != nil {
		return err
	}

	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}

	query := `
		INSERT INTO erajaya_product
			(
			 	id,
				name,
				price,
				description,
				quantity,
			 	created_at,
				updated_at
			)
		VALUES
			(
				$1,
				$2,
				$3,
				$4,
				$5,
				CURRENT_TIMESTAMP,
			 	CURRENT_TIMESTAMP
			)
	`

	_, err = tx.Exec(
		context.Background(),
		query,
		product.ID,
		product.Name,
		product.Price,
		product.Description,
		product.Quantity,
	)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (p *ProductRepository) GetListSort(ctx context.Context, product entity.ListProductRequest) ([]entity.Product, error) {
	db, err := p.client.GetSlave(ctx)
	if err != nil {
		return nil, err
	}

	query := `SELECT id,
				name,
				price,
				description,
				quantity,
			 	created_at,
				updated_at
			  FROM erajaya_product`

	if product.SortBy == entity.GetListSortByNewProduct {
		query += " ORDER BY created_at DESC"
	} else if product.SortBy == entity.GetListSortByCheapestPrice {
		query += " ORDER BY price ASC"
	} else if product.SortBy == entity.GetListSortByExpensivePrice {
		query += " ORDER BY price DESC"
	} else if product.SortBy == entity.GetListSortByProductNameASC {
		query += " ORDER BY name ASC"
	} else if product.SortBy == entity.GetListSortByProductNameDESC {
		query += " ORDER BY name DESC"
	}

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	var result []entity.Product
	for rows.Next() {
		var row entity.Product
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.Price,
			&row.Description,
			&row.Quantity,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, row)
	}

	return result, nil
}
