package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Webhook : A struct which respresent the webhook table
type Webhook struct {
	ID          uint `gorm:"primary_key" json:"id"`
	CityID      uint `json:"city_id"`
	City        City
	CallbackURL string `json:"callback_url"`
}

//Request : Sends a request to the webhook
func (w *Webhook) Request(t *Temperature) int {
	requestBody, err := json.Marshal(map[string]interface{}{
		"city_id":   t.CityID,
		"max":       t.Max,
		"min":       t.Min,
		"timestamp": t.CreatedAt,
	})

	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.Post(w.CallbackURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(res.StatusCode)
	return res.StatusCode
}
