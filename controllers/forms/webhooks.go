package forms

//Webhook : A struct that represent the webhook api
type Webhook struct {
	CityID      uint   `json:"city_id" binding:"required"`
	CallbackURL string `json:"callback_url" binding:"required"`
}
