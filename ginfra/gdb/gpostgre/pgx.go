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

func NewPgxPostgres(pool *pgxpool.Pool) *PostgresPgx {
	return &PostgresPgx{
		Commander: &pgxCommander{pool: pool},
		Builder:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		Pool:      pool,
	}
}

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
