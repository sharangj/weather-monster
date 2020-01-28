package forms

//City : Represents the form data for creating a city
type City struct {
	Name      string `json:"name" binding:"required"`
	Latitude  int    `json:"latitude" binding:"required"`
	Longitude int    `json:"longitude" binding:"required"`
}
