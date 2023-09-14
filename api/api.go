package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"twitter-back-end/controllers"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // always allow connections
	},
}

var clients = make(map[int]*websocket.Conn)

func HandleApi(r *gin.Engine) {

	r.POST("/createUser", controllers.CreateUser)

	r.GET("/getUser", controllers.GetUser)

	r.POST("/postTweet", controllers.AuthenticateJWT, controllers.PostTweet)

	r.GET("/getTweets", controllers.AuthenticateJWT, controllers.GetTweet)

	r.GET("/validateUser", controllers.AuthenticateJWT, controllers.ValidateUser)

	r.GET("/findOtherUsers", controllers.AuthenticateJWT, controllers.FindOtherUsers)

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

	r.PUT("updateUserPassword", controllers.AuthenticateJWT, controllers.UpdateUserPassword)

	r.GET("generateToken", controllers.GenerateToken)

	r.POST("sentMessage", controllers.AuthenticateJWT, controllers.SentMessage)

	r.GET("getMessages", controllers.AuthenticateJWT, controllers.GetMessages)

	r.GET("getConversations", controllers.AuthenticateJWT, controllers.GetConversations)

	r.GET("getStatus", controllers.AuthenticateJWT, controllers.GetStatus)

	r.GET("getOnlineStatus", controllers.AuthenticateJWT, controllers.GetOnlineStatus)

	r.PUT("updateStatus", controllers.UpdateStatus)

	r.GET("/echo", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("Failed to upgrade the connection: %v\n", err)
			return
		}

		var params struct {
			Id int
		}
		c.Bind(&params)

		clients[params.Id] = conn

		fmt.Println("\n", "Client Connected: ", params.Id)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			var receivedMsg structs.Message
			err = json.Unmarshal(msg, &receivedMsg)

			fmt.Println(receivedMsg)

			recipientConn, ok := clients[receivedMsg.RecieverId]
			if ok {
				if err = recipientConn.WriteJSON(receivedMsg); err != nil {
					fmt.Printf("Error sending message to %d: %v\n", receivedMsg.RecieverId, err)
				}
			}
		}

	})
}
