package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a heartbeat endpoint
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
