package models

//City : Declares the structure of City database model
type City struct {
	Model
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
