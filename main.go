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

	r.Use(cors.New(corsConfig))

	r.POST("/createUser", controllers.CreateUser)
	r.POST("/getUser", controllers.GetUser)
	r.POST("/postTweet", controllers.PostTweet)
	r.POST("/getTweets", controllers.GetTweet)
	r.POST("/validateUser", controllers.ValidateUser)
	r.POST("/getPeopleToFollow", controllers.GetPeopleToFollow)
	r.POST("/addtofollowerList", controllers.AddtofollowerList)
	r.POST("getFollowersTweet", controllers.GetFollowersTweet)
	r.POST("getFollowing", controllers.GetFollowing)
	r.POST("deleteFollower", controllers.DeleteFollower)
	r.POST("getFollowers", controllers.GetFollowers)
	r.POST("getIfTweetLiked", controllers.GetIfTweetLiked)
	r.POST("likeTweet", controllers.LikeTweet)
	r.POST("unlikeTweet", controllers.UnlikeTweet)
	r.POST("getLikesOnTweet", controllers.GetLikesOnTweet)
	r.POST("submitComment", controllers.SubmitComment)
	r.POST("showCommentsOnTweet", controllers.ShowCommentsOnTweet)
	r.POST("getTotalCommentOnTweet", controllers.GetTotalCommentOnTweet)
	r.PUT("updateTweetContent", controllers.UpdateTweetContent)
	r.POST("deleteTweet", controllers.DeleteTweet)
	r.POST("updateUserData", controllers.UpdateUserData)
	r.POST("addProfilePicture", controllers.AddProfilePicture)
	r.Run() // listen and serve on 0.0.0.0:8080
}
