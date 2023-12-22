package gpostgre

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type txPgx struct {
	pool *pgxpool.Pool
}

func NewTxPgx(pool *pgxpool.Pool) gdb.Tx {
	return &txPgx{
		pool: pool,
	}
}

func (t *txPgx) DoTransaction(ctx context.Context, opt *gdb.TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx pgx.Tx
	if opts.IsoLevel == "" {
		tx, err = t.pool.Begin(ctx)
	} else {
		tx, err = t.pool.BeginTx(ctx, opts)
	}
	if err != nil {
		return err
	}

	var commit bool
	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(gdb.ErrRollback, errRollback)
			}
			panic(p)
		} else if commit {
			if errCommit := tx.Commit(ctx); errCommit != nil {
				err = errors.Join(gdb.ErrCommit, errCommit)
			}
		} else if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(gdb.ErrRollback, errRollback)
			}
		} else {
			if errCommit := tx.Commit(ctx); errCommit != nil {
				err = errors.Join(gdb.ErrCommit, errCommit)
			}
		}
	}()

	txKey := context.WithValue(ctx, gdb.TxKey{}, tx)

	commit, err = fn(txKey)

	return err
}

func (t *txPgx) extractOpt(opt *gdb.TxOption) (opts pgx.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != gdb.TxTypePgx && opt.Type != gdb.TxTypeNone {
		err = fmt.Errorf("%w, your type is not *pgx.TxOptions. but %s", gdb.ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(pgx.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not pgx.TxOptions", gdb.ErrTypeTx)
		return
	}

	return opts, nil
}
