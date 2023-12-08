package gdb

import (
	"context"
)

type TxType uint

const (
	TxTypeNone TxType = iota
	TxTypeMongoDB
	TxTypePgx
	TxTypeSqlx
)

type TxOption struct {
	Type   TxType
	Option interface{}
}

type Tx interface {
	DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) error) (err error)
}
