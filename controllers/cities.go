package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharangj/weather_monster/controllers/forms"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
)

//CitiesController : A Controller for all city opertations
type CitiesController struct{}

//Create : Creates a city
func (h CitiesController) Create(c *gin.Context) {
	var cityParams forms.City
	db := db.Connect()

	err := c.BindJSON(&cityParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	city := models.City{Name: cityParams.Name, Latitude: cityParams.Latitude, Longitude: cityParams.Longitude}
	result := db.Create(&city)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, city)
	return
}
