package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/tests"
)

var _ = Describe("City", func() {
	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("creates a city", func() {
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			Expect(city.Name).To(Equal("New York"))
			Expect(city.Latitude).To(Equal(float64(42)))
			Expect(city.Longitude).To(Equal(float64(12)))
		})
	})
})
