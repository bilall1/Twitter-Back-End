package main

import (
	"twitter-back-end/controllers"
	"twitter-back-end/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {

	//initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	r := gin.Default()

	//r.Use(cors.Default())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization") // Add "Authorization" header
	corsConfig.AddAllowHeaders("ThirdParty")    // Add "Authorization" header

	r.Use(cors.New(corsConfig))

	r.POST("/createUser", controllers.CreateUser)
	r.GET("/getUser", controllers.GetUser)
	r.POST("/postTweet", controllers.AuthenticateJWT, controllers.PostTweet)

	r.GET("/getTweets", controllers.AuthenticateJWT, controllers.GetTweet)

	r.GET("/validateUser", controllers.AuthenticateJWT, controllers.ValidateUser)

	r.GET("/getPeopleToFollow", controllers.AuthenticateJWT, controllers.GetPeopleToFollow)

	r.POST("/addtofollowerList", controllers.AuthenticateJWT, controllers.AddtofollowerList)

	r.GET("getFollowersTweet", controllers.AuthenticateJWT, controllers.GetFollowersTweet)
	r.GET("getFollowing", controllers.AuthenticateJWT, controllers.GetFollowing)

	r.DELETE("deleteFollower", controllers.AuthenticateJWT, controllers.DeleteFollower)

	r.GET("getFollowers", controllers.AuthenticateJWT, controllers.GetFollowers)
	r.GET("getIfTweetLiked", controllers.AuthenticateJWT, controllers.GetIfTweetLiked)

	r.POST("likeTweet", controllers.AuthenticateJWT, controllers.LikeTweet)
	r.POST("unlikeTweet", controllers.AuthenticateJWT, controllers.UnlikeTweet)

	r.GET("getLikesOnTweet", controllers.AuthenticateJWT, controllers.GetLikesOnTweet)

	r.POST("submitComment", controllers.AuthenticateJWT, controllers.SubmitComment)

	r.GET("showCommentsOnTweet", controllers.AuthenticateJWT, controllers.ShowCommentsOnTweet)

	r.GET("getTotalCommentOnTweet", controllers.AuthenticateJWT, controllers.GetTotalCommentOnTweet)

	r.PUT("updateTweetContent", controllers.AuthenticateJWT, controllers.UpdateTweetContent)

	r.DELETE("deleteTweet", controllers.AuthenticateJWT, controllers.DeleteTweet)

	r.PUT("updateUserData", controllers.AuthenticateJWT, controllers.UpdateUserData)

	r.POST("addProfilePicture", controllers.AuthenticateJWT, controllers.AddProfilePicture)

	r.GET("getTotalFollowers", controllers.AuthenticateJWT, controllers.GetTotalFollowers)
	r.GET("getTotalFollowings", controllers.AuthenticateJWT, controllers.GetTotalFollowings)

	r.PUT("updateUserPassword", controllers.AuthenticateJWT, controllers.UpdateUserPassword) //put

	r.GET("generateToken", controllers.GenerateToken)

	r.Run() // listen and serve on 0.0.0.0:8080
}
