package gcommon

import (
	"context"
)

type CloseFn func(ctx context.Context) error
type CloseFnx func() error

// PanicIfError will trigger a panic if the provided error is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
