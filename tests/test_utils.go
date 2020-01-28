package tests

import "github.com/jinzhu/gorm"

//DbCleanup : Cleaning up the DB
func DbCleanup(db *gorm.DB) {
	db.Exec("TRUNCATE cities")
}
