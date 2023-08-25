package main

import (
	"github.com/bilall1/twitter-backend/controllers"
	"github.com/bilall1/twitter-backend/initializers"
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

	r.POST("/createUser", controllers.AuthenticateJWT, controllers.CreateUser)
	r.POST("/getUser", controllers.GetUser)
	r.POST("/postTweet", controllers.AuthenticateJWT, controllers.PostTweet)
	r.POST("/getTweets", controllers.AuthenticateJWT, controllers.GetTweet)
	r.POST("/validateUser", controllers.AuthenticateJWT, controllers.ValidateUser)
	r.POST("/getPeopleToFollow", controllers.AuthenticateJWT, controllers.GetPeopleToFollow)
	r.POST("/addtofollowerList", controllers.AuthenticateJWT, controllers.AddtofollowerList)
	r.POST("getFollowersTweet", controllers.AuthenticateJWT, controllers.GetFollowersTweet)
	r.POST("getFollowing", controllers.AuthenticateJWT, controllers.GetFollowing)
	r.POST("deleteFollower", controllers.AuthenticateJWT, controllers.DeleteFollower)
	r.POST("getFollowers", controllers.AuthenticateJWT, controllers.GetFollowers)
	r.POST("getIfTweetLiked", controllers.AuthenticateJWT, controllers.GetIfTweetLiked)
	r.POST("likeTweet", controllers.AuthenticateJWT, controllers.LikeTweet)
	r.POST("unlikeTweet", controllers.AuthenticateJWT, controllers.UnlikeTweet)
	r.POST("getLikesOnTweet", controllers.AuthenticateJWT, controllers.GetLikesOnTweet)
	r.POST("submitComment", controllers.AuthenticateJWT, controllers.SubmitComment)
	r.POST("showCommentsOnTweet", controllers.AuthenticateJWT, controllers.ShowCommentsOnTweet)
	r.POST("getTotalCommentOnTweet", controllers.AuthenticateJWT, controllers.GetTotalCommentOnTweet)
	r.PUT("updateTweetContent", controllers.AuthenticateJWT, controllers.UpdateTweetContent)
	r.POST("deleteTweet", controllers.AuthenticateJWT, controllers.DeleteTweet)
	r.POST("updateUserData", controllers.AuthenticateJWT, controllers.UpdateUserData)
	r.POST("addProfilePicture", controllers.AuthenticateJWT, controllers.AddProfilePicture)
	r.POST("getTotalFollowers", controllers.AuthenticateJWT, controllers.GetTotalFollowers)
	r.POST("getTotalFollowings", controllers.AuthenticateJWT, controllers.GetTotalFollowings)
	r.POST("updateUserPassword", controllers.AuthenticateJWT, controllers.UpdateUserPassword)
	r.POST("generateToken", controllers.GenerateToken)
	r.Run() // listen and serve on 0.0.0.0:8080
}
