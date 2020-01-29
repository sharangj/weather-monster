package controllers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/server"
	"github.com/sharangj/weather_monster/tests"
)

var _ = Describe("Forecasts", func() {

	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Get", func() {
		It("forecasts the average min temperature and max temperature for a city in the last 24 hours", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)
			temperature1 := models.Temperature{CityID: city.ID, Max: 30, Min: 25, CreatedAt: time.Now().Add(-10 * time.Hour)}
			temperature2 := models.Temperature{CityID: city.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-6 * time.Hour)}
			temperature3 := models.Temperature{CityID: city.ID, Max: 32, Min: 26, CreatedAt: time.Now().Add(-11 * time.Hour)}
			temperature4 := models.Temperature{CityID: city.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-30 * time.Hour)}

			db.Create(&temperature1)
			db.Create(&temperature2)
			db.Create(&temperature3)
			db.Create(&temperature4)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/forecasts/%d", city.ID), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["city_id"]).To(Equal(float64(city.ID)))
			Expect(response["max"]).To(Equal(float64(31)))
			Expect(response["min"]).To(Equal(float64(25)))
			Expect(response["sample"]).To(Equal(float64(3)))
		})

		It("forecasts the average min temperature and max temperature for a city in the last 24 hours when there are multiple cities", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			city2 := models.City{Name: "London", Latitude: 50, Longitude: 13}
			db.Create(&city2)
			temperature1 := models.Temperature{CityID: city.ID, Max: 30, Min: 25, CreatedAt: time.Now().Add(-10 * time.Hour)}
			temperature2 := models.Temperature{CityID: city.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-6 * time.Hour)}
			temperature3 := models.Temperature{CityID: city2.ID, Max: 32, Min: 26, CreatedAt: time.Now().Add(-11 * time.Hour)}
			temperature4 := models.Temperature{CityID: city.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-30 * time.Hour)}

			db.Create(&temperature1)
			db.Create(&temperature2)
			db.Create(&temperature3)
			db.Create(&temperature4)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/forecasts/%d", city.ID), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["city_id"]).To(Equal(float64(city.ID)))
			Expect(response["max"]).To(Equal(float64(30.5)))
			Expect(response["min"]).To(Equal(float64(24.5)))
			Expect(response["sample"]).To(Equal(float64(2)))
		})

		It("does not forecast when the city id is wrong", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/forecasts/%d", city.ID+1), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(response["error"]).To(Equal("City with this Id is not found"))
		})

		It("shows an empty response when there is no sample set", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			city2 := models.City{Name: "London", Latitude: 50, Longitude: 13}
			db.Create(&city2)
			temperature1 := models.Temperature{CityID: city.ID, Max: 30, Min: 25, CreatedAt: time.Now().Add(-10 * time.Hour)}
			temperature2 := models.Temperature{CityID: city.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-6 * time.Hour)}
			temperature3 := models.Temperature{CityID: city.ID, Max: 32, Min: 26, CreatedAt: time.Now().Add(-11 * time.Hour)}
			temperature4 := models.Temperature{CityID: city2.ID, Max: 31, Min: 24, CreatedAt: time.Now().Add(-30 * time.Hour)}

			db.Create(&temperature1)
			db.Create(&temperature2)
			db.Create(&temperature3)
			db.Create(&temperature4)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/forecasts/%d", city2.ID), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(response["error"]).To(Equal("record not found"))
		})
	})
})
