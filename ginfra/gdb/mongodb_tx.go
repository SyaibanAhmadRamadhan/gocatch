package gdb

import (
	"context"

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
	var session mongo.Session
	if opt == nil {
		session, err = m.client.StartSession()
	} else {
		opts := m.extractOpt(opt)
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

func (m *txMongodb) extractOpt(opt *TxOption) (opts *options.SessionOptions) {
	if opt == nil {
		return nil
	}

	if opt.Option == nil || opt.Type != TxTypeMongoDB {
		return nil
	}

	opts = opt.Option.(*options.SessionOptions)
	return opts
}
