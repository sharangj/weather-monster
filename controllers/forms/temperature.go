package forms

//Temperature :  A form to validate the params of TemperaturesController
type Temperature struct {
	CityID uint    `json:"city_id" binding:"required"`
	Max    float64 `json:"max" binding:"required"`
	Min    float64 `json:"min" binding:"required"`
}
