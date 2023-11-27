package JOpg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RDBMSPgx interface {
	QCount(ctx context.Context, sql string, args ...any) (count int64, err error)
	QAll(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error)
	BeginRun(ctx context.Context, fn func(rdbms RDBMSPgx) error) (err error)
	BeginTxRun(ctx context.Context, opts pgx.TxOptions, fn func(rdbms RDBMSPgx) error) (err error)
	PgxCommander
}

type rdbmsPgxImpl struct {
	conn *pgxpool.Pool
	PgxCommander
}

func NewRDBMS(conn *pgxpool.Pool) RDBMSPgx {
	return &rdbmsPgxImpl{
		conn:         conn,
		PgxCommander: conn,
	}
}

func (r *rdbmsPgxImpl) QCount(ctx context.Context, sql string, args ...any) (count int64, err error) {
	err = r.QueryRow(ctx, sql, args...).Scan(&count)
	return
}

func (r *rdbmsPgxImpl) QAll(ctx context.Context, sql string, args ...any) (rows pgx.Rows, err error) {
	return r.Query(ctx, sql, args...)
}

func (r *rdbmsPgxImpl) BeginRun(ctx context.Context, fn func(rdbms RDBMSPgx) error) (err error) {
	tx, err := r.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed start tx begin | err: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(ctx); err != nil {
				panic(err)
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				panic(err)
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				panic(err)
			}
		}
	}()

	err = fn(&rdbmsPgxImpl{
		PgxCommander: tx,
	})

	return err
}

func (r *rdbmsPgxImpl) BeginTxRun(ctx context.Context, opts pgx.TxOptions, fn func(rdbms RDBMSPgx) error) (err error) {
	tx, err := r.conn.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed start BeginTx | err: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(ctx); err != nil {
				panic(err)
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				panic(err)
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				panic(err)
			}
		}
	}()

	err = fn(&rdbmsPgxImpl{
		PgxCommander: tx,
	})

	return err
}
