package app

import (
	"github.com/mritunjaykumar/users-api/controllers/ping"
	"github.com/mritunjaykumar/users-api/controllers/users"
	"github.com/mritunjaykumar/users-api/logger"
)

func mapUrls() {
	// ping endpoint
	router.GET("/ping", ping.Ping)

	// users endpoints
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

	logger.Log("Mapped all URLs")
}
