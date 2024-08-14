package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"module/placeholder/internal/db/migrations"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	conn    *sql.DB
	closeFn func() error
}

func (d *DB) Conn() *sql.DB {
	return d.conn
}

func (d *DB) Close() error {
	if d.closeFn == nil {
		return nil
	}
	return d.closeFn()
}

func LocalPG(dsn string) (*DB, error) {
	sDb, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error making connection: %v", err)
	}
	return &DB{
		conn:    sDb,
		closeFn: sDb.Close,
	}, nil
}

func CloudSQL(dsn, instanceConnectionName string) (*DB, error) {
	d, err := cloudsqlconn.NewDialer(context.Background(), cloudsqlconn.WithIAMAuthN())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	opts = append(opts, cloudsqlconn.WithPrivateIP())

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName, opts...)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return &DB{
		conn: dbPool,
	}, nil
}

func MigrateTables(db *DB) error {
	if db == nil {
		log.Println("migrations: db is nil")
		return nil
	}
	_, err := db.Conn().Exec(migrations.Migrations)
	if err != nil {
		return fmt.Errorf("error executing sql: %v", err)
	}
	return nil
}
