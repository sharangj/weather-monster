package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharangj/weather_monster/controllers/forms"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
)

//WebhooksController : A struct to represent operations on webhooks
type WebhooksController struct{}

//Create : A function that creates a webook
func (h WebhooksController) Create(c *gin.Context) {
	var webhookParams forms.Webhook
	db := db.Connect()

	err := c.BindJSON(&webhookParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	var city models.City
	city.ID = webhookParams.CityID

	result := db.First(&city)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "The city was not found",
		})
		return
	}

	webhook := models.Webhook{CityID: webhookParams.CityID, CallbackURL: webhookParams.CallbackURL}
	result = db.Create(&webhook)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, webhook)
	return
}

//Delete : A function to delete a webhook
func (h WebhooksController) Delete(c *gin.Context) {
	db := db.Connect()

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "Please pass valid id"})
		return
	}

	var webhook models.Webhook
	webhook.ID = uint(ID)

	result := db.First(&webhook)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Webhook with this Id is not found",
		})
		return
	}

	result = db.Delete(&webhook)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, webhook)
	return
}
