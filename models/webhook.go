package models

//Webhook : A struct which respresent the webhook table
type Webhook struct {
	ID          uint `gorm:"primary_key" json:"id"`
	CityID      uint `json:"city_id"`
	City        City
	CallbackURL string `json:"callback_url"`
}
