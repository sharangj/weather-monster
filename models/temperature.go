package models

import "time"

//Temperature : A struct to represent the temperature table
type Temperature struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CityID    uint `json:"city_id"`
	City      City
	Max       float64   `json:"max"`
	Min       float64   `json:"min"`
	CreatedAt time.Time `json:"created_at"`
}
