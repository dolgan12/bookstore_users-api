package users

import (
	"net/http"
	"strconv"

	"github.com/dolgan12/bookstore_users-api/domain/users"
	"github.com/dolgan12/bookstore_users-api/services"
	"github.com/dolgan12/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// GetUser Gets a user
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("userid should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser Creates a user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// SearchUser Finds a user
//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
