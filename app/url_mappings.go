package app

import (
	"github.com/mritunjaykumar/users-api/controllers/ping"
	"github.com/mritunjaykumar/users-api/controllers/users"
	"github.com/mritunjaykumar/users-api/logger"
)

func mapUrls() {
	// ping endpoint
	r.GET("/ping", ping.Ping)

	// users endpoints
	r.GET("/users/:user_id", users.GetUser)
	r.POST("/users", users.CreateUser)

	logger.Log("Mapped all URLs")
}
