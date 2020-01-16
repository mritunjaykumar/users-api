package logger

import (
	"bytes"
	"net/http"

	"github.com/mritunjaykumar/users-api/util/errors"
)

const (
	influxDBWriteEndPoint = "http://localhost:8186/write"
	contentType           = "text/plain"
)

// WriteMetrics writes metrics to the telegraf
func WriteMetrics(body []byte) error {
	resp, err := http.Post(influxDBWriteEndPoint, contentType, bytes.NewBuffer(body))
	if err != nil {
		Error("writeMetrics failed.",
			errors.RestErr{
				Message: "writeMetrics failed.",
				Status:  500,
				Err:     err.Error(),
			})

		return err
	}

	defer resp.Body.Close()
	return nil
}
