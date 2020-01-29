package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

var _ = Describe("Webhooks", func() {

	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("creates a webhook for a city", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"city_id":      city.ID,
				"callback_url": "https://www.foobar.com/temp_hook",
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["city_id"]).To(Equal(float64(city.ID)))
			Expect(response["callback_url"]).To(Equal("https://www.foobar.com/temp_hook"))
		})

		It("does not create a webhook when the city id is wrong", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"city_id":      city.ID + 1,
				"callback_url": "https://www.foobar.com/temp_hook",
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(requestBody))
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

	Describe("Delete", func() {
		It("deletes the specified webhook", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			webhook := models.Webhook{CityID: city.ID, CallbackURL: "https://www.foobar.com/hook"}
			db.Create(&webhook)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/webhooks/%d", webhook.ID), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["city_id"]).To(Equal(float64(city.ID)))
			Expect(response["callback_url"]).To(Equal(webhook.CallbackURL))
		})
	})
})
