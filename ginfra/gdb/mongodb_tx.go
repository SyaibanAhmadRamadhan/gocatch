package gdb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type txMongodb struct {
	client *mongo.Client
}

func NewTxMongodb(client *mongo.Client) Tx {
	return &txMongodb{client: client}
}

func (m *txMongodb) DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) error) (err error) {
	opts, err := m.extractOpt(opt)
	if err != nil {
		return
	}

	var session mongo.Session
	if opts == nil {
		session, err = m.client.StartSession()
	} else {
		session, err = m.client.StartSession(opts)
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				if err = session.AbortTransaction(sc); err != nil {
					return
				}
				panic(p)
			} else if err != nil {
				if err = session.AbortTransaction(sc); err != nil {
					return
				}
			} else {
				if err = session.CommitTransaction(sc); err != nil {
					return
				}
			}
		}()

		return fn(sc)
	})

	return err
}

func (m *txMongodb) extractOpt(opt *TxOption) (opts *options.SessionOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != TxTypeMongoDB && opt.Type != TxTypeNone {
		err = fmt.Errorf("%w, your type is not pgx. but %s", ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*options.SessionOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *options.SessionOptions", ErrTypeTx)
		return
	}

	return opts, nil
}
