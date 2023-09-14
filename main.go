package main

import (
	"twitter-back-end/initializers"

	"twitter-back-end/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {

	//initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization") // Add "Authorization" header
	corsConfig.AddAllowHeaders("ThirdParty")    // Add "Authorization" header

	r.Use(cors.New(corsConfig))

	api.HandleApi(r)

	r.Run()
}
