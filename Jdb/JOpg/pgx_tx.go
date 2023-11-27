package JOpg

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx interface {
	BeginRun(ctx context.Context, fn func(tx pgx.Tx) error) (err error)
	BeginTxRun(ctx context.Context, opts pgx.TxOptions, fn func(tx pgx.Tx) error) (err error)
}
