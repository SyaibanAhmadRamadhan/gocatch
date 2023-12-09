package gdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txPgx struct {
	pool *pgxpool.Pool
}

func NewTxPgx(pool *pgxpool.Pool) Tx {
	return &txPgx{
		pool: pool,
	}
}

func (t *txPgx) DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) error) (err error) {
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

	defer func() {
		if p := recover(); p != nil {
			if err = tx.Rollback(ctx); err != nil {
				return
			}
			panic(p)
		} else if err != nil {
			if err = tx.Rollback(ctx); err != nil {
				return
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				return
			}
		}
	}()

	txKey := context.WithValue(ctx, TxKey{}, tx)

	err = fn(txKey)

	return err
}

func (t *txPgx) extractOpt(opt *TxOption) (opts pgx.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != TxTypePgx && opt.Type != TxTypeNone {
		err = fmt.Errorf("%w, your type is not *pgx.TxOptions. but %s", ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(pgx.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not pgx.TxOptions", ErrTypeTx)
		return
	}

	return opts, nil
}
