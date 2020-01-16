package app

import "time"

type ginLog struct {
	Timestamp    string        `json:"timestamp"`
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Protocol     string        `json:"protocol"`
	Status       int           `json:"status"`
	Latency      time.Duration `json:"latency"`
	UserAgent    string        `json:"userAgent"`
	ErrorMessage string        `json:"errorMessage"`
}
