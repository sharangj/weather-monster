package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/tests"
)

var _ = Describe("Temperature", func() {
	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("creates a temperature", func() {
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			temperature1 := models.Temperature{CityID: city.ID, Max: 30, Min: 25}
			db.Create(&temperature1)

			Expect(temperature1.CityID).To(Equal(city.ID))
		})
	})
})
