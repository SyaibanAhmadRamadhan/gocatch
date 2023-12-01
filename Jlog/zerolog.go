package Jlog

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitZeroLogDev(sample *zerolog.BurstSampler) {
	buildInfo, _ := debug.ReadBuildInfo()

	if sample != nil {
		logger := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", buildInfo.GoVersion).
			Logger().
			Sample(sample)
		log.Logger = logger
	} else {
		logger := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", buildInfo.GoVersion).
			Logger()
		log.Logger = logger
	}

	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func InitZerologFile(cusLog *zerolog.Logger, filePath string) {
	file, err := os.OpenFile(
		filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Err(err).Msg("failed closed file")
		}
	}(file)

	if cusLog != nil {
		log.Logger = *cusLog
	} else {
		logger := zerolog.New(file).With().Timestamp().Logger()
		log.Logger = logger
	}

	log.Info().Msg("successfully init logger")

}

func InitZerolog(hook *ZerologHook, level zerolog.Level) {
	if hook != nil {
		logger := zerolog.New(nil).Level(level).With().Timestamp().Logger().Hook(hook)
		log.Logger = logger
	} else {
		logger := zerolog.New(os.Stderr).Level(level).With().Timestamp().Logger()
		log.Logger = logger
	}

	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")

	hook.Wait()
}
