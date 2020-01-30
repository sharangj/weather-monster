package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Temperature : A struct to represent the temperature table
type Temperature struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CityID    uint `json:"city_id"`
	City      City
	Max       float64   `json:"max"`
	Min       float64   `json:"min"`
	CreatedAt time.Time `json:"created_at"`
}

//AfterCreate : Function called after creating the webhook
func (t *Temperature) AfterCreate(scope *gorm.Scope) (err error) {
	if t.ID > 0 {
		var webhooks []Webhook

		scope.DB().Where("city_id = ?", t.CityID).Find(&webhooks)
		for _, webhook := range webhooks {
			go webhook.Request(t)
		}
	}
	return
}
