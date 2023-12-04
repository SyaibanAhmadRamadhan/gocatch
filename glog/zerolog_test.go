package glog

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
)

func TestZerolog(t *testing.T) {
	ctx := context.Background()

	c := NewLogCtx(ctx)

	l := Ctx(c)
	test2(c)

	l.Logger.Info().Msg("output")
}

func test2(ctx context.Context) {
	l := Ctx(ctx)

	l.Err(l.Str("test", "test").Str("test2", "test2"))

	// test1(ctx)
}

func test1(ctx context.Context) {
	l := Ctx(ctx)

	l.Err(zerolog.Dict().Str("test", "test"))
}
