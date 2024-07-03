package gormpg

import (
	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PG(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func CloudSQL(dsn string) (*gorm.DB, error) {
	_, err := pgxv5.RegisterDriver("cloudsql-postgres",
		cloudsqlconn.WithIAMAuthN(),
		cloudsqlconn.WithDefaultDialOptions(
			cloudsqlconn.WithPrivateIP(),
		))
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsql-postgres",
		DSN:        dsn,
	}))
}
