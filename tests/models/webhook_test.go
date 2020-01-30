package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/tests"
)

var _ = Describe("Webhook", func() {
	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("sends a request to the webhook URL", func() {
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)
			webhook := models.Webhook{CityID: city.ID, CallbackURL: "https://www.foobar.com/hook"}
			db.Create(&webhook)

			Expect(webhook.CityID).To(Equal(city.ID))
			Expect(webhook.CallbackURL).To(Equal("https://www.foobar.com/hook"))
		})
	})
})
