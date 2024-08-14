package migrations

import (
	_ "embed"
)

//go:embed migrations.sql
var Migrations string
