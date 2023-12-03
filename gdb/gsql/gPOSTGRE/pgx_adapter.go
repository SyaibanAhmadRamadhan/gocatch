package gPOSTGRE

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenPgxPool(connString string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := pgxpool.New(ctx, connString)
	gcommon.PanicIfError(err)

	err = conn.Ping(ctx)
	gcommon.PanicIfError(err)

	return conn
}

func OpenPgxConn(connString string, withPing bool) *pgx.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, connString)
	gcommon.PanicIfError(err)

	if withPing {
		err = conn.Ping(ctx)
		gcommon.PanicIfError(err)
	}

	return conn
}
