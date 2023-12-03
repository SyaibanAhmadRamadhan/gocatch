package gPOSTGRE

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenConnSqlxPgx(connStr string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sqlx.Connect("pgx", connStr)

	err = db.PingContext(ctx)
	gcommon.PanicIfError(err)

	return db
}
