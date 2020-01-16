package logger

import (
	"github.com/mritunjaykumar/users-api/util/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	// Logger creates zap logger based on the config
	logger *zap.Logger
)

const (
	utcTimeFormat = "2006-01-02T15:04:05Z0700"
)

// UTC time encode
func utcTimeEncode(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(utcTimeFormat))
}

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = utcTimeEncode
	zapConfig.EncoderConfig.TimeKey = "timestamp"

	var err error
	if logger, err = zapConfig.Build(); err != nil {
		panic(err)
	}
}

// Log wraps zap "Info" function
func Log(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
	logger.Sync()
}

// Debug wraps zap "Debug" function
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
	logger.Sync()
}

// Error wraps zap "Error" function
func Error(msg string, err errors.RestErr, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))

	logger.Error(msg, fields...)
	logger.Sync()
}
