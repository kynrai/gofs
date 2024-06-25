package pgx

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// New pgx connection pool with a DSN if the format of:
// postgres://username:password@localhost:5432/database_name
// returns a pool of connections. This is primarily a conevience function
// to set some sane defaults for the connection pool. For more advanced usage,
// use the pgxpool package directly.
func New(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return pgxpool.NewWithConfig(context.Background(), config)
}
