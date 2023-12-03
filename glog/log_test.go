package glog

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
)

func TestLog(t *testing.T) {
	ctx := NewLogCtx(context.Background())
	l := zerolog.Ctx(ctx)
	AppendLogCtx(ctx, func(c zerolog.Context) zerolog.Context {
		return c.Str("test", "test")
	})

	service(ctx)

	l.Info().Msg("output")
}

func Repo(ctx context.Context) {
	logMsg := NewLogMsg(KeyRepoLayer)
	defer func() {
		AppendLogCtxWithMsg(ctx, logMsg)
	}()

	logMsg.Msg = &Msg{
		Level:  "info",
		Msg:    "rama",
		Caller: CallerInfoStr(),
	}

}
func service(ctx context.Context) {
	AppendLogCtx(ctx, func(c zerolog.Context) zerolog.Context {
		return c.Str("service2", "searchQuery").Str("service", "pageNum")
	})

	Repo(ctx)
}
