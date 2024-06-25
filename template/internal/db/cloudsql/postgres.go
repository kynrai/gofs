package main

import (
	"context"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// New creates a new pgx connection pool with a DSN and instance connection name
//
// dsn format "user=myuser password=mypass dbname=mydb sslmode=disable"
//
// icn format "project:region:instance"
//
// returns a pool of connections. This is primarily a conevience function.
// returns the pool, a cleanup function, and an error. The cleanup function
// should be called when the pool is no longer needed to close the connection
func New(ctx context.Context, Context, dsn, icn string) (*pgxpool.Pool, func() error, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, nil, err
	}

	d, err := cloudsqlconn.NewDialer(
		ctx,
		cloudsqlconn.WithIAMAuthN(),
		cloudsqlconn.WithDefaultDialOptions(
			cloudsqlconn.WithPrivateIP(),
		),
	)
	if err != nil {
		return nil, nil, err
	}

	config.ConnConfig.DialFunc = func(ctx context.Context, _ string, instance string) (net.Conn, error) {
		return d.Dial(ctx, icn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, nil, err
	}
	cleanupFn := func() error { return d.Close() }

	return pool, cleanupFn, nil
}
