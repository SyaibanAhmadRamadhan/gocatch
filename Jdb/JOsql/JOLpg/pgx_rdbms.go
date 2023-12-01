package JOLpg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOsql"
)

type rdbmsPgxImpl struct {
	conn *pgxpool.Pool
	PgxCommander
}

func NewRDBMSpgx(conn *pgxpool.Pool) JOsql.RDBMS {
	return &rdbmsPgxImpl{
		conn:         conn,
		PgxCommander: conn,
	}
}

func (r *rdbmsPgxImpl) Write(ctx context.Context, sql string, namedArg map[string]any) (rowsAffected int64, err error) {
	if namedArg != nil {
		args := pgx.NamedArgs(namedArg)
		res, err := r.Exec(ctx, sql, args)
		rowsAffected = res.RowsAffected()

		if err != nil {
			return 0, err
		}

	} else {
		res, err := r.Exec(ctx, sql)
		rowsAffected = res.RowsAffected()

		if err != nil {
			return 0, err
		}
	}

	return

}

func (r *rdbmsPgxImpl) QueryCount(ctx context.Context, sql string, namedArg map[string]any) (count int64, err error) {
	if namedArg != nil {
		args := pgx.NamedArgs(namedArg)
		err = r.QueryRow(ctx, sql, args).Scan(&count)
	} else {
		err = r.QueryRow(ctx, sql).Scan(&count)
	}
	return
}

func (r *rdbmsPgxImpl) CheckOne(ctx context.Context, sql string, namedArg map[string]any) (b bool, err error) {
	if namedArg != nil {
		args := pgx.NamedArgs(namedArg)
		err = r.QueryRow(ctx, sql, args).Scan(&b)
	} else {
		err = r.QueryRow(ctx, sql).Scan(&b)
	}
	return
}

func (r *rdbmsPgxImpl) QueryAll(ctx context.Context, sql string, namedArg map[string]any) (results []map[string]any, err error) {
	var rows pgx.Rows
	if namedArg != nil {
		args := pgx.NamedArgs(namedArg)
		rows, err = r.Query(ctx, sql, args)
	} else {
		rows, err = r.Query(ctx, sql)
	}

	if err != nil {
		return
	}

	var fields []string
	for _, v := range rows.FieldDescriptions() {
		fields = append(fields, v.Name)
	}

	results = make([]map[string]any, 0)

	for rows.Next() {
		result := make(map[string]any)
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		if len(values) != len(fields) {
			return nil, fmt.Errorf("query all data, but values and fields not the same | fields: %v, values: %v", fields, values)
		}

		for i, v := range values {
			result[fields[i]] = v
		}

		results = append(results, result)
	}

	return
}

func (r *rdbmsPgxImpl) QueryOne(ctx context.Context, sql string, namedArg map[string]any) (result map[string]any, err error) {
	var rows pgx.Rows
	if namedArg != nil {
		args := pgx.NamedArgs(namedArg)
		rows, err = r.Query(ctx, sql, args)
	} else {
		rows, err = r.Query(ctx, sql)
	}

	if err != nil {
		return
	}

	var fields []string
	for _, v := range rows.FieldDescriptions() {
		fields = append(fields, v.Name)
	}

	for rows.Next() {
		result = make(map[string]any)
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		if len(values) != len(fields) {
			return nil, fmt.Errorf("query all data, but values and fields not the same | fields: %v, values: %v", fields, values)
		}

		for i, v := range values {
			result[fields[i]] = v
		}

		break
	}

	return
}

func (r *rdbmsPgxImpl) BeginRun(ctx context.Context, fn func(rdbms JOsql.RDBMS) error) (err error) {
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

func (r *rdbmsPgxImpl) BeginTxRun(ctx context.Context, opts JOsql.TxOptions, fn func(rdbms JOsql.RDBMS) error) (err error) {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.TxIsoLevel(opts.IsoLevel),
		AccessMode:     pgx.TxAccessMode(opts.AccessMode),
		DeferrableMode: pgx.TxDeferrableMode(opts.DeferrableMode),
		BeginQuery:     opts.BeginQuery,
	})
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