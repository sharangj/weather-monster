package controllers

import (
	"net/http"
	"strconv"

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

//Update : A function to update the attributes of a city
func (h CitiesController) Update(c *gin.Context) {

	var cityParams forms.UpdateCity
	db := db.Connect()

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "Please pass valid id"})
		return
	}

	err = c.BindJSON(&cityParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	var city models.City
	city.ID = uint(ID)

	result := db.First(&city)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusBadRequest,
			"error": "City with this Id is not found",
		})
		return
	}

	result = db.Model(&city).Updates(&cityParams)

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

//Delete : A function to update the attributes of a city
func (h CitiesController) Delete(c *gin.Context) {
	db := db.Connect()

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "Please pass valid id"})
		return
	}

	var city models.City
	city.ID = uint(ID)

	result := db.First(&city)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusBadRequest,
			"error": "City with this Id is not found",
		})
		return
	}

	result = db.Delete(&city)

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
