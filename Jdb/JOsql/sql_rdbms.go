package JOsql

import (
	"context"
)

type RDBMS interface {
	Write(ctx context.Context, sql string, namedArg map[string]any) (rowsAffected int64, err error)

	QueryCount(ctx context.Context, sql string, namedArg map[string]any) (count int64, err error)
	QueryAll(ctx context.Context, sql string, namedArg map[string]any) (results []map[string]any, err error)
	QueryOne(ctx context.Context, sql string, namedArg map[string]any) (result map[string]any, err error)

	// CheckOne return true if data exist, otherwise return false. If error not nil, return false and error
	// query: SELECT EXISTS(SELECT 1 FROM table WHERE condition)
	CheckOne(ctx context.Context, sql string, namedArg map[string]any) (b bool, err error)

	BeginRun(ctx context.Context, fn func(rdbms RDBMS) error) (err error)
	BeginTxRun(ctx context.Context, opts TxOptions, fn func(rdbms RDBMS) error) (err error)
}
