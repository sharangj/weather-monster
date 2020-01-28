package models

//City : Declares the structure of City database model
type City struct {
	Model
	Name      string `json:"name"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}
