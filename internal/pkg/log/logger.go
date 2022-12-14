package log

import (
	"context"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
)

func init() {
	globalLogger = New(defaultLevel, os.Stdout)
}

func New(level zapcore.LevelEnabler, writer io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(writer),
			level,
		),
		options...,
	).Sugar()
}

func Logger(_ context.Context) *zap.SugaredLogger {
	return globalLogger
}

func StdLogger(_ context.Context) *log.Logger {
	return zap.NewStdLog(globalLogger.Desugar())
}
