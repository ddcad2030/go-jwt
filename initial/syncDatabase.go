package initial

import (
	"david/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
