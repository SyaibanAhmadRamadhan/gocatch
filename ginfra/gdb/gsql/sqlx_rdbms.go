package gsql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type Commander interface {
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type sqlxCommander struct {
	db *sqlx.DB
}

func (c *sqlxCommander) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.QueryxContext(ctx, query, args...)
	}

	return c.db.QueryxContext(ctx, query, args...)
}

func (c *sqlxCommander) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return c.db.ExecContext(ctx, query, args...)
}

func (c *sqlxCommander) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.NamedQuery(query, arg)
	}

	return c.db.NamedQueryContext(ctx, query, arg)
}

func (c *sqlxCommander) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.NamedExecContext(ctx, query, arg)
	}

	return c.db.NamedExecContext(ctx, query, arg)
}

func (c *sqlxCommander) extractTx(ctx context.Context) (tx *sqlx.Tx) {
	val := ctx.Value(gdb.TxKey{})

	if tx, ok := val.(*sqlx.Tx); tx != nil && ok {
		return tx
	}

	return nil
}
