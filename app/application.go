package app

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
	"github.com/mritunjaykumar/users-api/util/errors"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	r = gin.New()
)

const (
	utcTimeFormat      = "2006-01-02T15:04:05Z0700"
	ginMeasurementName = "gin"
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

	metrics := fmt.Sprintf("%s,timestamp=%s,method=%s,path=%s,protocol=%s,status=%s latency=%d",
		ginMeasurementName, param.TimeStamp.UTC().Format(utcTimeFormat),
		param.Method, param.Path, param.Request.Proto, strconv.Itoa(param.StatusCode), param.Latency,
	)

	// Already logged error if this error out, so not checking for error here
	logger.WriteMetrics([]byte(metrics))

	var bSlice []byte
	var err error
	bSlice, err = json.Marshal(glog)

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
