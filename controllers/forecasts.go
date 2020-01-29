package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
)

//ForecastsController : A struct to represent operations on forecasts
type ForecastsController struct{}

//Get : A function which returns the average min and max temperatures for the last 24 hours
func (h ForecastsController) Get(c *gin.Context) {
	db := db.Connect()
	CityID, err := strconv.Atoi(c.Param("city_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "Please pass valid city id"})
		return
	}

	var city models.City
	city.ID = uint(CityID)

	result := db.First(&city)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusBadRequest,
			"error": "City with this Id is not found",
		})
		return
	}

	var forecasts models.Forecast
	result = db.Table("temperatures").Select("city_id, AVG(max) as max, AVG(min) as min, count(*) as sample").Where("created_at > now() - interval '24 hours'").Group("city_id").Having("city_id = ?", city.ID).Scan(&forecasts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, forecasts)
	return
}
