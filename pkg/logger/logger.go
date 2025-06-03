package logger

import (
	"context"

	"github.com/zuhrulumam/go-parking-lot/pkg/ctxkeys"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json", // use "console" for human-readable logs
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "xtime",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	zap.ReplaceGlobals(logger) // Optional: makes logger globally available
	return logger
}

func LogWithCtx(ctx context.Context, logger *zap.Logger, msg string, fields ...zap.Field) {
	logger.With(
		zap.String("correlation_id", ctx.Value(ctxkeys.CtxKeyCorrelationID).(string)),
		zap.String("path", ctx.Value(ctxkeys.CtxKeyPath).(string)),
		zap.String("method", ctx.Value(ctxkeys.CtxKeyMethod).(string)),
	).Info(msg, fields...)
}
