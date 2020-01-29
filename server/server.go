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
	temperature := new(controllers.TemperaturesController)
	forecasts := new(controllers.ForecastsController)
	webhooks := new(controllers.WebhooksController)

	r.GET("/ping", ping.Status)
	r.POST("/cities", city.Create)
	r.PATCH("/cities/:id", city.Update)
	r.DELETE("/cities/:id", city.Delete)
	r.POST("/temperatures", temperature.Create)
	r.GET("/forecasts/:city_id", forecasts.Get)
	r.POST("/webhooks", webhooks.Create)
	r.DELETE("/webhooks/:id", webhooks.Delete)

	return r
}
