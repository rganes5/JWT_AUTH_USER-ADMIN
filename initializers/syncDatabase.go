package initializers

import "github.com/rganes5/go-jwt-auth/models"

// Automatically create database tables based on the struct definition of a model.
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
