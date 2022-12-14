package log

import (
	"context"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Desugar() *zap.Logger
	Named(name string) *zap.SugaredLogger
	WithOptions(opts ...zap.Option) *zap.SugaredLogger
	With(args ...interface{}) *zap.SugaredLogger
	Level() zapcore.Level
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	DPanicln(args ...interface{})
	Panicln(args ...interface{})
	Fatalln(args ...interface{})
	Sync() error
}

var (
	globalLogger Logger
	DefaultLevel = zap.DebugLevel
)

func init() {
	globalLogger = New(DefaultLevel, os.Stdout)
}

func New(level zapcore.LevelEnabler, writer io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = DefaultLevel
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

func SetLogger(logger Logger) {
	globalLogger = logger
}

func L(_ context.Context) Logger {
	return globalLogger
}

func StdLogger(_ context.Context) *log.Logger {
	return zap.NewStdLog(globalLogger.Desugar())
}
