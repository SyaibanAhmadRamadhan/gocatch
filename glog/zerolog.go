package glog

import (
	"context"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/SyaibanAhmadRamadhan/gocatch/gtime"
)

var once sync.Once
var log zerolog.Logger

// initLog function initializes a zerolog.Logger object once
// which can be reused across the application to perform logging.
func initLog() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel) // default to INFO
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if os.Getenv("APP_ENV") != "development" {
			// 	TODO: add log prod
		}

		var gitRevision string

		buildInfo, _ := debug.ReadBuildInfo()

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Caller().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
		zerolog.DefaultContextLogger = &log
	})

	return log
}

// NewLogCtx function creates and returns a new context with the logger from initLog function.
func NewLogCtx(ctx context.Context) context.Context {
	lCtx := initLog().WithContext(ctx)
	zerolog.Ctx(lCtx).UpdateContext(func(ctx zerolog.Context) zerolog.Context {
		return ctx.Str("time", time.Now().Format(time.RFC3339))
	})
	return lCtx
}

type ctx struct {
	start  time.Time
	Logger *zerolog.Logger
	context.Context
	caller *CallInfo
}

func Ctx(c context.Context) *ctx {
	return &ctx{
		Context: c,
		start:   time.Now(),
		Logger:  zerolog.Ctx(c),
		caller:  CallerInfo(2),
	}
}

const keyExecutionTime = "execution_time"
const keyFileName = "filename"
const keyPackageName = "package"

func (c *ctx) with(fields *zerolog.Event) {
	c.Logger.UpdateContext(func(ctx zerolog.Context) zerolog.Context {
		executionTime := gtime.TimeTrack(c.start)
		return ctx.Dict(c.caller.FuncName, fields.
			Str(keyExecutionTime, executionTime).
			Str(keyFileName, c.caller.FileName).
			Str(keyPackageName, c.caller.PackageName),
		)
	})
}
func (c *ctx) Dict() *zerolog.Event {
	return zerolog.Dict()
}

func (c *ctx) Str(key string, val string) *zerolog.Event {
	return zerolog.Dict().Str(key, val)
}

func (c *ctx) Info(fields *zerolog.Event) {
	c.with(fields.Str("level", "info"))
}

func (c *ctx) Err(fields *zerolog.Event) {
	c.with(fields.Str("level", "error"))
}
