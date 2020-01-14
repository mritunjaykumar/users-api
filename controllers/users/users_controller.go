package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mritunjaykumar/users-api/domain/users"
	"github.com/mritunjaykumar/users-api/logger"
	"github.com/mritunjaykumar/users-api/services"
	"github.com/mritunjaykumar/users-api/util/errors"
)

// CreateUser creates user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("invalid json body")
		logger.Error(restErr.Message, *restErr)
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		logger.Error(saveErr.Message, *saveErr)
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	logger.Log(fmt.Sprintf("Created user %s %s with email address %s",
		result.FirstName, result.LastName, result.Email))
}

// GetUser gets user
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		restErr := errors.BadRequest("invalid user id")
		logger.Error(restErr.Message, *restErr)
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		logger.Error(getErr.Message, *getErr)
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
	logger.Log(fmt.Sprintf("Got the user with userID [%d]", userID))
}
