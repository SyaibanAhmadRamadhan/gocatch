package gmongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type txMongodb struct {
	client *mongo.Client
}

func NewTxMongodb(client *mongo.Client) gdb.Tx {
	return &txMongodb{client: client}
}

func (m *txMongodb) DoTransaction(ctx context.Context, opt *gdb.TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
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

		var commit bool
		defer func() {
			if p := recover(); p != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					return
				}
				panic(p)
			} else if commit {
				if errCommit := session.CommitTransaction(sc); errCommit != nil {
					return
				}
			} else if err != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					return
				}
			} else {
				if errCommit := session.CommitTransaction(sc); errCommit != nil {
					return
				}
			}
		}()

		commit, err = fn(sc)
		return err
	})

	return err
}

func (m *txMongodb) extractOpt(opt *gdb.TxOption) (opts *options.SessionOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != gdb.TxTypeMongoDB && opt.Type != gdb.TxTypeNone {
		err = fmt.Errorf("%w, your type is not pgx. but %s", gdb.ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*options.SessionOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *options.SessionOptions", gdb.ErrTypeTx)
		return
	}

	return opts, nil
}
