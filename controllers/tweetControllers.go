package controllers

import (
	"twitter-back-end/models"
	"twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
)

func PostTweet(c *gin.Context) {
	var body models.Tweet
	c.Bind(&body)

	tweetResponse, err := services.PostTweet(body.Content, body.Id, body.Link)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Tweet": tweetResponse,
	})

}

func GetTweet(c *gin.Context) {
	var params structs.TweetOfUser
	c.Bind(&params)

	tweetsResponse, err := services.GetTweet(params.Email, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	if len(tweetsResponse) == 0 {
		c.JSON(200, gin.H{
			"Tweets": []models.Tweet{},
		})
		return
	}
	c.JSON(200, gin.H{
		"Tweets": tweetsResponse,
	})
}

func GetFollowersTweet(c *gin.Context) {
	var params structs.TweetFollower
	c.Bind(&params)

	tweetsResponse, err := services.GetFollowersTweet(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Tweets": tweetsResponse,
	})

}

func GetIfTweetLiked(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	isLikedResponse, err := services.GetIfTweetLiked(params.TweetId, params.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": isLikedResponse,
	})

}

func LikeTweet(c *gin.Context) {
	var body structs.TweetUser
	c.Bind(&body)

	setLikedResponse, err := services.LikeTweet(body.TweetId, body.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": setLikedResponse,
	})

}

func UnlikeTweet(c *gin.Context) {
	var body structs.TweetUser
	c.Bind(&body)

	setUnLikedResponse, err := services.UnLikeTweet(body.TweetId, body.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": setUnLikedResponse,
	})

}

func GetLikesOnTweet(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	totalLikesResponse, err := services.GetLikesOnTweet(params.TweetId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": totalLikesResponse,
	})

}

func SubmitComment(c *gin.Context) {
	var body structs.TweetComment
	c.Bind(&body)

	setCommentResponse, err := services.SubmitComment(body.TweetId, body.UserId, body.Content)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Comment": setCommentResponse,
	})

}

func ShowCommentsOnTweet(c *gin.Context) {
	var params structs.CommentLimit
	c.Bind(&params)

	allCommentsResponse, err := services.ShowCommentsOnTweet(params.TweetId, params.Limit)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Comments": allCommentsResponse,
	})

}

func GetTotalCommentOnTweet(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	totalCommentsResponse, err := services.GetCommentsOnTweet(params.TweetId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": totalCommentsResponse,
	})

}

func UpdateTweetContent(c *gin.Context) {
	var params models.Tweet
	c.Bind(&params)

	updatedResponse, err := services.UpdateContent(params.Id, params.Content)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": updatedResponse,
	})

}

func DeleteTweet(c *gin.Context) {
	var body models.Tweet
	c.Bind(&body)

	deletedResponse, err := services.DeleteTweet(body.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Deleted": deletedResponse,
	})

}
