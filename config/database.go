package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var PGConnection *bun.DB

func CreateDatabaseConnection() {
	dsn := "postgres://postgres:postgres@localhost:5432/employee_app?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	PGConnection = db
}
