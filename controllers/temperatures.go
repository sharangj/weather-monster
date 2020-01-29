package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharangj/weather_monster/controllers/forms"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
)

//TemperaturesController : A struct to initialize the controller
type TemperaturesController struct{}

//Create : A method to create temperatures for a city
func (h TemperaturesController) Create(c *gin.Context) {
	var temperatureParams forms.Temperature
	db := db.Connect()

	err := c.BindJSON(&temperatureParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	if temperatureParams.Max < temperatureParams.Min {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Max temperature cannot be lesser than Min",
		})
		return
	}

	var city models.City
	city.ID = temperatureParams.CityID

	result := db.First(&city)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusBadRequest,
			"error": "The city was not found",
		})
		return
	}

	temperature := models.Temperature{CityID: temperatureParams.CityID, Max: temperatureParams.Max, Min: temperatureParams.Min}
	result = db.Create(&temperature)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, temperature)
	return
}
