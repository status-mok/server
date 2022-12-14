package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type noOPLogger struct {
	logger *zap.Logger
}

func NewNOOPLogger() *noOPLogger {
	return &noOPLogger{
		logger: zap.NewNop(),
	}
}

func (l noOPLogger) Desugar() *zap.Logger { return l.logger }

func (l noOPLogger) Named(_ string) *zap.SugaredLogger { return l.logger.Sugar() }

func (l noOPLogger) WithOptions(_ ...zap.Option) *zap.SugaredLogger { return l.logger.Sugar() }

func (l noOPLogger) With(_ ...interface{}) *zap.SugaredLogger { return l.logger.Sugar() }

func (l noOPLogger) Level() zapcore.Level { return l.logger.Level() }

func (l noOPLogger) Debug(_ ...interface{}) {}

func (l noOPLogger) Info(_ ...interface{}) {}

func (l noOPLogger) Warn(_ ...interface{}) {}

func (l noOPLogger) Error(_ ...interface{}) {}

func (l noOPLogger) DPanic(_ ...interface{}) {}

func (l noOPLogger) Panic(_ ...interface{}) {}

func (l noOPLogger) Fatal(_ ...interface{}) {}

func (l noOPLogger) Debugf(_ string, _ ...interface{}) {}

func (l noOPLogger) Infof(_ string, _ ...interface{}) {}

func (l noOPLogger) Warnf(_ string, _ ...interface{}) {}

func (l noOPLogger) Errorf(_ string, _ ...interface{}) {}

func (l noOPLogger) DPanicf(_ string, _ ...interface{}) {}

func (l noOPLogger) Panicf(_ string, _ ...interface{}) {}

func (l noOPLogger) Fatalf(_ string, _ ...interface{}) {}

func (l noOPLogger) Debugw(_ string, _ ...interface{}) {}

func (l noOPLogger) Infow(_ string, _ ...interface{}) {}

func (l noOPLogger) Warnw(_ string, _ ...interface{}) {}

func (l noOPLogger) Errorw(_ string, _ ...interface{}) {}

func (l noOPLogger) DPanicw(_ string, _ ...interface{}) {}

func (l noOPLogger) Panicw(_ string, _ ...interface{}) {}

func (l noOPLogger) Fatalw(_ string, _ ...interface{}) {}

func (l noOPLogger) Debugln(_ ...interface{}) {}

func (l noOPLogger) Infoln(_ ...interface{}) {}

func (l noOPLogger) Warnln(_ ...interface{}) {}

func (l noOPLogger) Errorln(_ ...interface{}) {}

func (l noOPLogger) DPanicln(_ ...interface{}) {}

func (l noOPLogger) Panicln(_ ...interface{}) {}

func (l noOPLogger) Fatalln(_ ...interface{}) {}

func (l noOPLogger) Sync() error { return nil }
