package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once          sync.Once
	logger        *zap.Logger
	encoderConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "line",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logConfig = zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    10000,
			Thereafter: 10,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
)

func Logger() *zap.Logger {
	return zapLogger()
}

func zapLogger() *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = logConfig.Build()
		if err != nil {
			panic(err)
		}
		callerOpt := zap.AddCallerSkip(1)
		logger = logger.WithOptions(callerOpt)
	})
	return logger
}
