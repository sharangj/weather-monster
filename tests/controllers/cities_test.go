package controllers_test

import (
	"bytes"
	"encoding/json"
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
