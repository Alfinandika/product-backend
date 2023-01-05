package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgreClientItf is a client to interact with PostgreSQL database using pgx.
type PostgreClientItf interface {
	// GetMaster returns connection a read write client to postgres.
	GetMaster(ctx context.Context) (*pgxpool.Pool, error)

	// GetSlave returns connection a read only client to postgres.
	GetSlave(ctx context.Context) (*pgxpool.Pool, error)

	// Close closes all connection to the postgres.
	Close()
}
