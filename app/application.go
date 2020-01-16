package app

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
	"github.com/mritunjaykumar/users-api/util/errors"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	r = gin.New()
)

const (
	utcTimeFormat = "2006-01-02T15:04:05Z0700"
)

func customLogFormatter(param gin.LogFormatterParams) string {
	glog := &ginLog{
		Timestamp:    param.TimeStamp.UTC().Format(utcTimeFormat),
		Method:       param.Method,
		Path:         param.Path,
		Protocol:     param.Request.Proto,
		Status:       param.StatusCode,
		Latency:      param.Latency,
		UserAgent:    param.Request.UserAgent(),
		ErrorMessage: param.ErrorMessage,
	}

	bSlice, err := json.Marshal(glog)

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

// StartApplication starts the application
func StartApplication() {
	logger.Log("Starting application...")

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.Use(gin.LoggerWithFormatter(customLogFormatter))
	r.Use(gin.Recovery())

	mapUrls()
	r.Run(":8080")
}
