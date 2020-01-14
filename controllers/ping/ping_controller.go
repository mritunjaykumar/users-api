package ping

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/logger"
	"net/http"
)

// Ping is a health check function
func Ping(c *gin.Context) {
	logger.Log(fmt.Sprintf("Request URI [%s]", c.Request.RequestURI))
	c.String(http.StatusOK, "Pong")
}
