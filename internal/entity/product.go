package entity

import "time"

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Price       int       `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Quantity    int       `json:"quantity" db:"quantity"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ListProductRequest struct {
	SortBy string `json:"sort_by"`
}
