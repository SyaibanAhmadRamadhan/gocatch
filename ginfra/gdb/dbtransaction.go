package gdb

import (
	"context"
	"errors"
)

type TxType uint

const (
	TxTypeNone TxType = iota
	TxTypeMongoDB
	TxTypePgx
	TxTypeSqlx
)

func (t TxType) String() string {
	switch t {
	case TxTypeNone:
		return "none type"
	case TxTypeMongoDB:
		return "mongodb"
	case TxTypePgx:
		return "pgx"
	case TxTypeSqlx:
		return "sqlx"
	}

	return ""
}

var ErrTypeTx = errors.New("invalid type tx")

type TxOption struct {
	Type   TxType
	Option interface{}
}

type TxKey struct{}

type Tx interface {
	DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) error) (err error)
}
