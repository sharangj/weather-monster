package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PingController : Controller format to check if server is up
type PingController struct{}

//Status : Return a simple message to the browser
func (h PingController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
