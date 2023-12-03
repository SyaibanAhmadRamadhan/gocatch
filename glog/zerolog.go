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

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
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

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

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

// NewLog function returns a new logger by calling initLog function.
func NewLog() zerolog.Logger {
	return initLog()
}

// NewLogCtx function creates and returns a new context with the logger from initLog function.
func NewLogCtx(ctx context.Context) context.Context {
	lCtx := initLog().WithContext(ctx)
	return lCtx
}

// AppendLogCtx function appends the provided function to the logger in the context.
func AppendLogCtx(ctx context.Context, fn func(zerolog.Context) zerolog.Context) {
	lCtx := zerolog.Ctx(ctx)
	lCtx.UpdateContext(fn)
}

// AppendLogCtxWithMsg function appends the provided message to the logger in the context.
func AppendLogCtxWithMsg(ctx context.Context, msg logMessage) {
	AppendLogCtx(ctx, func(c zerolog.Context) zerolog.Context {
		dict := zerolog.Dict()

		data := gcommon.Ternary(msg.Data == nil, nil, msg.Data)
		dict.Interface("data", data)

		message := gcommon.Ternary(msg.Msg == nil, nil, msg.Msg)
		dict.Interface("message", message)

		return c.Dict(
			msg.key, dict.
				Str(KeyExecutionTime, gtime.TimeTrack(msg.startTime)),
		)
	})
}

// func InitZeroLogDev(sample *zerolog.BurstSampler) {
// 	buildInfo, _ := debug.ReadBuildInfo()
//
// 	if sample != nil {
// 		logger := zerolog.New(zerolog.ConsoleWriter{
// 			Out:        os.Stdout,
// 			TimeFormat: time.RFC3339,
// 		}).
// 			Level(zerolog.TraceLevel).
// 			With().
// 			Timestamp().
// 			Caller().
// 			Int("pid", os.Getpid()).
// 			Str("go_version", buildInfo.GoVersion).
// 			Logger().
// 			Sample(sample)
// 		log.Logger = logger
// 	} else {
// 		logger := zerolog.New(zerolog.ConsoleWriter{
// 			Out:        os.Stdout,
// 			TimeFormat: time.RFC3339,
// 		}).
// 			Level(zerolog.TraceLevel).
// 			With().
// 			Timestamp().
// 			Caller().
// 			Int("pid", os.Getpid()).
// 			Str("go_version", buildInfo.GoVersion).
// 			Logger()
// 		log.Logger = logger
// 	}
//
// 	log.Trace().Msg("trace message")
// 	log.Debug().Msg("debug message")
// 	log.Info().Msg("info message")
// 	log.Warn().Msg("warn message")
// 	log.Error().Msg("error message")
// 	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
// 	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
// }
//
// func InitZerologFile(cusLog *zerolog.Logger, filePath string) {
// 	file, err := os.OpenFile(
// 		filePath,
// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
// 		0664,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func(file *os.File) {
// 		err := file.Close()
// 		if err != nil {
// 			log.Err(err).Msg("failed closed file")
// 		}
// 	}(file)
//
// 	if cusLog != nil {
// 		log.Logger = *cusLog
// 	} else {
// 		logger := zerolog.New(file).With().Timestamp().Logger()
// 		log.Logger = logger
// 	}
//
// 	log.Info().Msg("successfully init logger")
//
// }
//
// func InitZerolog(hook *ZerologHook, level zerolog.Level) {
// 	if hook != nil {
// 		logger := zerolog.New(nil).Level(level).With().Timestamp().Logger().Hook(hook)
// 		log.Logger = logger
// 	} else {
// 		logger := zerolog.New(os.Stderr).Level(level).With().Timestamp().Logger()
// 		log.Logger = logger
// 	}
//
// 	log.Trace().Msg("trace message")
// 	log.Debug().Msg("debug message")
// 	log.Info().Msg("info message")
// 	log.Warn().Msg("warn message")
// 	log.Error().Msg("error message")
// 	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
// 	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
//
// 	hook.Wait()
// }
