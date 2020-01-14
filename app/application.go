package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
)

var (
	router = gin.Default()
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")

	// Pass on to the next-in-chain
	c.Next()

	fmt.Println("Out of dummy")
}

// StartApplication starts the application
func StartApplication() {
	logger.Log("Starting application...")

	router.Use(DummyMiddleware)

	mapUrls()
	router.Run(":8080")
}
