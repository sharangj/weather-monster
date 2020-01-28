package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/sharangj/weather_monster/db"
	"github.com/sharangj/weather_monster/models"
	"github.com/sharangj/weather_monster/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sharangj/weather_monster/server"
)

var _ = Describe("Cities", func() {

	AfterEach(func() {
		db := db.Connect()
		tests.DbCleanup(db)
	})

	Describe("Create", func() {
		It("creates a city", func() {
			router := server.Init()
			var response map[string]interface{}

			requestBody, err := json.Marshal(map[string]interface{}{
				"name":      "New York",
				"latitude":  52,
				"longitude": 13,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/cities", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["name"]).To(Equal("New York"))
		})

		It("does not create a city when the city is already created", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"name":      "New York",
				"latitude":  52,
				"longitude": 13,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/cities", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(response["error"]).ToNot(BeNil())
		})

		It("does not create a city when wrong params are passed", func() {
			router := server.Init()
			var response map[string]interface{}

			requestBody, err := json.Marshal(map[string]interface{}{
				"name": "New York",
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/cities", bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(response["error"]).ToNot(BeNil())
		})
	})

	Describe("PATCH", func() {
		It("updates a city when the right params are passed to it", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"latitude":  52,
				"longitude": 13,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/cities/%d", city.ID), bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["name"]).To(Equal(city.Name))
			Expect(response["latitude"]).To(Equal(float64(52)))
			Expect(response["longitude"]).To(Equal(float64(13)))
		})

		It("updates a city when the only name is passed", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"name": "London",
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/cities/%d", city.ID), bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["name"]).To(Equal("London"))
			Expect(response["latitude"]).To(Equal(float64(42)))
			Expect(response["longitude"]).To(Equal(float64(12)))
		})

		It("does not update a city when the wrong id is passed", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			requestBody, err := json.Marshal(map[string]interface{}{
				"latitude":  52,
				"longitude": 13,
			})

			if err != nil {
				log.Fatalln(err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("/cities/%d", city.ID+1), bytes.NewBuffer(requestBody))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			err = json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(response["error"]).To(Equal("City with this Id is not found"))
		})
	})

	Describe("DELETE", func() {
		It("deletes the city when an id is passed", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/cities/%d", city.ID), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(response["name"]).To(Equal(city.Name))
			Expect(response["latitude"]).To(Equal(float64(city.Latitude)))
			Expect(response["longitude"]).To(Equal(float64(city.Longitude)))
		})

		It("does not delete a city when the wrong id is passed", func() {
			router := server.Init()
			var response map[string]interface{}

			//Create the city before the api call is made
			db := db.Connect()
			city := models.City{Name: "New York", Latitude: 42, Longitude: 12}
			db.Create(&city)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/cities/%d", city.ID+1), nil)
			router.ServeHTTP(w, req)

			err := json.Unmarshal([]byte(w.Body.String()), &response)

			if err != nil {
				log.Fatalln(err)
			}

			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(response["error"]).To(Equal("City with this Id is not found"))
		})
	})
})
