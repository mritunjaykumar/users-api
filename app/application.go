package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
	"github.com/mritunjaykumar/users-api/util/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

var (
	r = gin.New()
)

const (
	utcTimeFormat      = "2006-01-02T15:04:05Z0700"
	ginMeasurementName = "gin"
)

func customLogFormatter(param gin.LogFormatterParams) string {
	param.Keys = logger.GetGlobalTags()
	logMessage := make(map[string]interface{})
	logMessage["timestamp"] = param.TimeStamp.UTC().Format(utcTimeFormat)
	logMessage["method"] = param.Method
	logMessage["path"] = param.Path
	logMessage["protocol"] = param.Request.Proto
	logMessage["status"] = param.StatusCode
	logMessage["latency"] = param.Latency
	logMessage["userAgent"] = param.Request.UserAgent()
	logMessage["error"] = param.ErrorMessage

	for k,v := range param.Keys {
		logMessage[k] = fmt.Sprintf("%v", v)
	}

	//metrics := fmt.Sprintf("%s,method=%s,path=%s,protocol=%s,status=%s latency=%d %v",
	//	ginMeasurementName, param.Method, param.Path, param.Request.Proto,
	//	strconv.Itoa(param.StatusCode), param.Latency, param.TimeStamp.UnixNano(),
	//)
	//
	//// Already logged error if this error out, so not checking for error here
	//logger.WriteMetrics([]byte(metrics))

	var bSlice []byte
	var err error
	bSlice, err = json.Marshal(logMessage)

	if err != nil {
		logger.Error("customLogFormatter failed.",
			errors.RestErr{
				Message: "customLogFormatter failed.",
				Status:  500,
				Err:     err.Error(),
			})
	}

	return string(bSlice) + "\n"
}


func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		logValues := []zap.Field {
			zap.String("logger-context", "http-request-logger"),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("protocol", c.Request.Proto),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client-ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("start-time", start.Format(utcTimeFormat)),
			zap.String("end-time", end.Format(utcTimeFormat)),
			zap.String("latency-unit", "ns"),
			zap.Int64("latency", latency.Nanoseconds()),
		}

		if len(c.Errors) > 0 {
			var temp []string
			for _, err := range c.Errors.Errors() {
				temp = append(temp, err)
			}
			allErrorString := strings.Join(temp, ";")
			logger.Error("Error found",
				errors.RestErr{Message: allErrorString, Status: c.Writer.Status(), Err: ""}, logValues...)
		} else {
			logger.Log("Successfully processed", logValues...)
		}
	}
}

// StartApplication starts the application
func StartApplication() {
	logger.Log("Starting application...")

	//r.Use(gin.LoggerWithFormatter(customLogFormatter))
	r.Use(ZapLogger())
	r.Use(gin.Recovery())

	mapUrls()
	r.Run(":8095")
}
