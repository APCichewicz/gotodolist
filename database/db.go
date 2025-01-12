package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BunDb struct {
	db *bun.DB
}

func new(dsn string) BunDb {
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(pgdb, pgdialect.New())
	return BunDb{
		db: db,
	}
}

func (bdb *BunDb) get() *bun.DB {
	return bdb.db
}
