package gdb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type txSqlx struct {
	db *sqlx.DB
}

func NewTxSqlx(db *sqlx.DB) Tx {
	return &txSqlx{
		db: db,
	}
}

func (t *txSqlx) DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) error) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx *sqlx.Tx
	if opts == nil {
		tx, err = t.db.Beginx()
	} else {
		tx, err = t.db.BeginTxx(ctx, opts)
	}
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if err = tx.Rollback(); err != nil {
				return
			}
			panic(p)
		} else if err != nil {
			if err = tx.Rollback(); err != nil {
				return
			}
		} else {
			if err = tx.Commit(); err != nil {
				return
			}
		}
	}()

	txKey := context.WithValue(ctx, TxKey{}, tx)
	err = fn(txKey)

	return err
}

func (t *txSqlx) extractOpt(opt *TxOption) (opts *sql.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != TxTypeSqlx && opt.Type != TxTypeNone {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions. but %s", ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*sql.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions", ErrTypeTx)
		return
	}

	return opts, nil
}
