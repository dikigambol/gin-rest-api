package initializers

import "gin-rest-api/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
