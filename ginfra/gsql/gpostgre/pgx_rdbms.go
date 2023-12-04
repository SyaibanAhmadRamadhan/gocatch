package gpostgre

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Commander interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type RDBMS interface {
	BeginRun(ctx context.Context, fn func(tx Commander) error) error
	BeginTxRun(ctx context.Context, opts pgx.TxOptions, fn func(tx Commander) error) error
	Commander
}

type rdbmsImpl struct {
	pool *pgxpool.Pool
	Commander
}

func NewRdbmsPgx(pgxPool *pgxpool.Pool) RDBMS {
	return &rdbmsImpl{
		pool:      pgxPool,
		Commander: pgxPool,
	}
}

func (t *rdbmsImpl) BeginRun(ctx context.Context, fn func(tx Commander) error) (err error) {
	tx, err := t.Commander.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed start tx begin | err: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(ctx); err != nil {
				// TODO log error
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				// TODO log error
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				// TODO log error
			}
		}
	}()

	err = fn(tx)

	return err
}

func (t *rdbmsImpl) BeginTxRun(ctx context.Context, opts pgx.TxOptions, fn func(tx Commander) error) (err error) {
	tx, err := t.pool.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed start tx begin | err: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(ctx); err != nil {
				// TODO log error
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				// TODO log error
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				// TODO log error
			}
		}
	}()

	err = fn(tx)

	return err
}
