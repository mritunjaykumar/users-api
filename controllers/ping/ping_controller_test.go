package ping

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder() // Response recorder
	c, _ := gin.CreateTestContext(w)

	Ping(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Pong", w.Body.String())
}
