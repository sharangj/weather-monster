package models

//Forecast : A struct to define the model of forecast
type Forecast struct {
	CityID uint    `json:"city_id"`
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
	Sample float64 `json:"sample"`
}
