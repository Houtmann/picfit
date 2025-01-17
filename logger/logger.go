package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type ObjectEncoder = zapcore.ObjectEncoder

func New(cfg Config) *zap.Logger {
	var logger *zap.Logger
	if cfg.GetLevel() == DevelopmentLevel {
		logger, _ = NewDevelopmentLogger()
	} else {
		logger, _ = NewProductionLogger()
	}

	return logger
}

func String(k, v string) Field {
	return zap.String(k, v)
}

func Duration(k string, d time.Duration) Field {
	return zap.Duration(k, d)
}

func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

func Time(key string, val time.Time) Field {
	return zap.Time(key, val)
}

func Int(k string, i int) Field {
	return zap.Int(k, i)
}

func Array(key string, val zapcore.ArrayMarshaler) Field {
	return zap.Array(key, val)
}

func Int64(k string, i int64) Field {
	return zap.Int64(k, i)
}

func Error(v error) Field {
	return zap.Error(v)
}

func Object(key string, val zapcore.ObjectMarshaler) Field {
	return zap.Object(key, val)
}

func NewProductionLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewDevelopmentLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewNopLogger() (*zap.Logger, error) {
	return zap.NewNop(), nil
}
