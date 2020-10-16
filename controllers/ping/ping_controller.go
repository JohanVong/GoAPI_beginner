package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping function to test the ability to take traffic
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "true", "message": "pong"})
	return
}
