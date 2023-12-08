package gsql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Commander interface {
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type RDBMS interface {
	BeginTxRun(ctx context.Context, opts *sql.TxOptions, fn func(t Commander) error) error
	Commander
}

type commanderImpl struct {
	db *sqlx.DB
}

func (c *commanderImpl) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	stmt, err := c.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			// TODO log error
		}
	}()

	return stmt.QueryxContext(ctx, args...)
}

func (c *commanderImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			// TODO log error
		}
	}()

	return c.db.ExecContext(ctx, query, args...)
}

func (c *commanderImpl) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	stmt, err := c.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			// TODO log error
		}
	}()

	return stmt.Queryx(arg)
}

func (c *commanderImpl) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	stmt, err := c.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			// TODO log error
		}
	}()

	return stmt.ExecContext(ctx, arg)
}

type sqlxRdbms struct {
	db *sqlx.DB
	Commander
}

func NewSqlxCommander(db *sqlx.DB) RDBMS {
	return &sqlxRdbms{
		Commander: &commanderImpl{
			db: db,
		},
		db: db,
	}
}

func (s *sqlxRdbms) BeginTxRun(ctx context.Context, opts *sql.TxOptions, fn func(t Commander) error) error {
	tx, err := s.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				// TODO log error
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				// TODO log error
			}
		} else {
			if err = tx.Commit(); err != nil {
				// TODO log error
			}
		}
	}()

	err = fn(tx)

	return err
}
