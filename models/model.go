package models

import "time"

//Model : Gorm Model overwritten to accomodate json
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
