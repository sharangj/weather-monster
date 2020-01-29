package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/server"
	"github.com/sharangj/weather_monster/tests"
)

var _ = Describe("Temperatures", func() {

	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("creates a temperature", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"city_id": city.ID,
				"max":     30,
				"min":     25,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["city_id"]).To(Equal(float64(city.ID)))
			Expect(response["max"]).To(Equal(float64(30)))
			Expect(response["min"]).To(Equal(float64(25)))
		})

		It("does not create a temperature when max is less than min", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"city_id": city.ID,
				"max":     24,
				"min":     25,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(response["error"]).To(Equal("Max temperature cannot be lesser than Min"))
		})

		It("does not create a temperature when an invalid city id is passed", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"city_id": city.ID + 1,
				"max":     30,
				"min":     25,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(response["error"]).To(Equal("The city was not found"))
		})
	})
})
