package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type PostgreClient struct {
	master *pgxpool.Pool
	slaves []*pgxpool.Pool
}

// GetMaster returns master connection client to postgres.
func (p *PostgreClient) GetMaster(_ context.Context) (*pgxpool.Pool, error) {

	if p.master == nil {
		return nil, errors.New("connection not found")
	}
	return p.master, nil
}

// GetSlave returns slave connection client to postgres.
func (p *PostgreClient) GetSlave(_ context.Context) (*pgxpool.Pool, error) {

	if len(p.slaves) == 0 {
		return nil, errors.New("connection not found")
	}
	return p.slaves[0], nil
}

// Close closes all connection to the postgres.
func (p *PostgreClient) Close() {
	p.master.Close()

	for k := range p.slaves {
		p.slaves[k].Close()
	}
}

// NewPostgreClient creates a postgre client using pgx.
func NewPostgreClient() (*PostgreClient, error) {

	ctx := context.Background()

	// Connect master.
	master, err := connect(ctx, &dbConfig{
		Address:               "user=unicorn_user password=magical_password dbname=erajaya-database host=127.0.0.1 port=5432 sslmode=disable",
		MinConnection:         0,
		MaxConnection:         16,
		MaxConnectionLifetime: 600,
		MaxConnectionIdleTime: 300,
	})
	if err != nil {
		return nil, err
	}

	// Connect slaves.
	var slaves []*pgxpool.Pool
	conn, err := connect(ctx, &dbConfig{
		Address:               "user=unicorn_user password=magical_password dbname=erajaya-database host=127.0.0.1 port=5432 sslmode=disable",
		MinConnection:         0,
		MaxConnection:         16,
		MaxConnectionLifetime: 600,
		MaxConnectionIdleTime: 300,
	})
	if err != nil {
		return nil, err
	}

	slaves = append(slaves, conn)

	return &PostgreClient{
		master: master,
		slaves: slaves,
	}, nil
}

type dbConfig struct {
	Address               string
	MinConnection         int32
	MaxConnection         int32
	MaxConnectionLifetime int64
	MaxConnectionIdleTime int64
}

func connect(ctx context.Context, config *dbConfig) (*pgxpool.Pool, error) {
	// Parse configuration.
	dbConfig, err := pgxpool.ParseConfig(config.Address)
	if err != nil {
		return nil, err
	}

	dbConfig.MaxConns = config.MaxConnection
	dbConfig.MinConns = config.MinConnection
	dbConfig.MaxConnLifetime = func() time.Duration {
		ttl := config.MaxConnectionLifetime
		return time.Duration(ttl) * time.Second
	}()
	dbConfig.MaxConnIdleTime = func() time.Duration {
		ttl := config.MaxConnectionIdleTime
		return time.Duration(ttl) * time.Second
	}()

	// Create database connection.
	conn, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
