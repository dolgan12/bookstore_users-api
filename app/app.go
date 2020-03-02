package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is the startpoint of the application that sets up routing.
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
