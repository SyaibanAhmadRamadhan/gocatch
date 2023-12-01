package JOsql

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/SyaibanAhmadRamadhan/jolly"
)

func NewConnectionSqlx(uri string, driver string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := sqlx.MustConnect(driver, uri)
	err := db.PingContext(ctx)
	jolly.PanicIF(err)

	return db
}
