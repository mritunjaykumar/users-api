package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
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

	// ADD additional custom tags to the logs
	zapConfig.InitialFields = GetGlobalTags()
	zapConfig.Sampling = nil

	var err error
	if logger, err = zapConfig.Build(); err != nil {
		panic(err)
	}
}

// GetGlobalTags provides global tags added to the logs
func GetGlobalTags() map[string]interface{} {
	// ADD additional custom tags to the logs
	globalTags := make(map[string]interface{})
	globalTags["application"] = "astra"

	tempComponent := os.Args[0] // this might provide value like "/go/bin/usersapi"

	// Get just the app name and not the whole path. For example: out of "/go/bin/usersapi", just get "usersapi"
	globalTags["component"] = tempComponent[strings.LastIndex(tempComponent, "/")+1:]

	return globalTags
}

// Log wraps zap "Info" function
func Log(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
	logger.Sync()
}

// Debug wraps zap "Debug" function
func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
	logger.Sync()
}

// Error wraps zap "Error" function
func Error(message string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))

	logger.Error(message, fields...)
	logger.Sync()
}
