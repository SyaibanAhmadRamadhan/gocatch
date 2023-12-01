package Jlog

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestLog(t *testing.T) {
	// InitZeroLogDev(nil)

	// InitZerolog(&ZerologHook{
	// 	Notify:    "test",
	// 	Level:     zerolog.ErrorLevel,
	// 	WaitGroup: sync.WaitGroup{},
	// }, zerolog.ErrorLevel)

	ctx := context.Background()

	zerolog.New(os.Stdin).With().Ctx(ctx)

	f1(ctx)
	f2(ctx)

}

func f1(ctx context.Context) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("hello")

}
func f2(ctx context.Context) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("hello2")
}
