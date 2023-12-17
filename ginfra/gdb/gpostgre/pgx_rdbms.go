package gpostgre

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type Commander interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type pgxCommander struct {
	pool *pgxpool.Pool
}

func (r *pgxCommander) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Exec(ctx, sql, arguments...)
	}

	return r.pool.Exec(ctx, sql, arguments...)
}

func (r *pgxCommander) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Query(ctx, sql, args...)
	}
	return r.pool.Query(ctx, sql, args...)
}

func (r *pgxCommander) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *pgxCommander) Begin(ctx context.Context) (pgx.Tx, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Begin(ctx)
	}
	return r.Begin(ctx)
}

func (r *pgxCommander) extractTx(ctx context.Context) pgx.Tx {
	val := ctx.Value(gdb.TxKey{})

	if tx, ok := val.(pgx.Tx); tx != nil && ok {
		return tx
	}

	return nil
}
