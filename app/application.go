package app

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	router = gin.Default()
)

// LatencyMiddleware captures api latency
func LatencyMiddleware(c *gin.Context) {
	start := time.Now()

	// Pass on to the next-in-chain
	c.Next()

	elapsed := time.Since(start).Microseconds()
	logger.Log(fmt.Sprintf("Elapsed time in micro seconds: [%d]", elapsed))
}

// StartApplication starts the application
func StartApplication() {
	logger.Log("Starting application...")

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.Use(LatencyMiddleware)

	mapUrls()
	router.Run(":8080")
}
