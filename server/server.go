package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sharangj/weather_monster/controllers"
)

//Init : Initialise the Gin server
func Init() *gin.Engine {
	r := gin.Default()
	ping := new(controllers.PingController)
	city := new(controllers.CitiesController)

	r.GET("/ping", ping.Status)
	r.POST("/cities", city.Create)

	return r
}
