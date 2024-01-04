package gpostgre

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func OpenPgxPool(connString string, config ...*pgxpool.Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var err error
	var conn *pgxpool.Pool

	if config != nil && len(config) > 0 {
		conn, err = pgxpool.NewWithConfig(ctx, config[0])
	} else {
		conn, err = pgxpool.New(ctx, connString)
	}
	gcommon.PanicIfError(err)

	err = conn.Ping(ctx)
	gcommon.PanicIfError(err)

	return conn
}

func OpenPgxPoolWithConfig(config *pgxpool.Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	conn, err := pgxpool.NewWithConfig(ctx, config)
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
