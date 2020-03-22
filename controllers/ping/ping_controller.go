package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a health check function
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
