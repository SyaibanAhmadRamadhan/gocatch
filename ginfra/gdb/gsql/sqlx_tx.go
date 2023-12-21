package gsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type txSqlx struct {
	db *sqlx.DB
}

func NewTxSqlx(db *sqlx.DB) gdb.Tx {
	return &txSqlx{
		db: db,
	}
}

func (t *txSqlx) DoTransaction(ctx context.Context, opt *gdb.TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
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

	var commit bool

	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				return
			}
			panic(p)
		} else if commit {
			if errCommit := tx.Commit(); errCommit != nil {
				return
			}
		} else if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				return
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				return
			}
		}
	}()

	txKey := context.WithValue(ctx, gdb.TxKey{}, tx)

	commit, err = fn(txKey)

	return err
}

func (t *txSqlx) extractOpt(opt *gdb.TxOption) (opts *sql.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != gdb.TxTypeSqlx && opt.Type != gdb.TxTypeNone {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions. but %s", gdb.ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*sql.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions", gdb.ErrTypeTx)
		return
	}

	return opts, nil
}
