package forms

//City : Represents the form data for creating a city
type City struct {
	Name      string `json:"name" binding:"required"`
	Latitude  int    `json:"latitude" binding:"required"`
	Longitude int    `json:"longitude" binding:"required"`
}

//UpdateCity : Represents the form data for creating a city
type UpdateCity struct {
	Name      string `json:"name"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}
