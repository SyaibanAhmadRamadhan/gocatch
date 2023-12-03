package gPOSTGRE

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenConnSqlxPq(connString string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := sqlx.MustConnect("postgres", connString)

	err := db.PingContext(ctx)
	gcommon.PanicIfError(err)

	return db
}
