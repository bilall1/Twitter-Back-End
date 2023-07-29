package main

import (
	"github.com/bilall1/twitter-backend/initializers"
	"github.com/bilall1/twitter-backend/models"
)

func init() {

	// initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
