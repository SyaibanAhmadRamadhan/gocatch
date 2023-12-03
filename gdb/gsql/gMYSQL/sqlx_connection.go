package gMYSQL

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenConnSqlx(connString string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := sqlx.MustConnect("mysql", connString)

	err := db.PingContext(ctx)
	gcommon.PanicIfError(err)

	return db
}
