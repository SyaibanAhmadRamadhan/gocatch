package gpostgre

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

type PostgresPgx struct {
	Commander Commander
	Builder   squirrel.StatementBuilderType
	Pool      *pgxpool.Pool
}

func (p *PostgresPgx) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func NewPgxPostgres(connString string) *PostgresPgx {
	pgxPool := OpenPgxPool(connString)
	return &PostgresPgx{
		Commander: &pgxCommander{pool: pgxPool},
		Builder:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		Pool:      pgxPool,
	}
}

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
