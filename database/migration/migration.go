package migration

import (
	"learn-middleware/models"

	"gorm.io/gorm"
)

func MigrateUser(db *gorm.DB) {

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

}
