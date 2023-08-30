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

	tweet, err := services.PostTweet(body.Content, body.Id, body.Link)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Tweet": tweet,
	})

}

func GetTweet(c *gin.Context) {
	var params structs.TweetOfUser
	c.Bind(&params)

	tweets, err := services.GetTweet(params.Email, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	if len(tweets) == 0 {
		c.JSON(200, gin.H{
			"Tweets": []models.Tweet{},
		})
		return
	}
	c.JSON(200, gin.H{
		"Tweets": tweets,
	})
}

func GetFollowersTweet(c *gin.Context) {
	var params structs.TweetFollower
	c.Bind(&params)

	tweets, err := services.GetFollowersTweet(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Tweets": tweets,
	})

}

func GetIfTweetLiked(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	isLiked, err := services.GetIfTweetLiked(params.TweetId, params.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": isLiked,
	})

}

func LikeTweet(c *gin.Context) {
	var body structs.TweetUser
	c.Bind(&body)

	setLiked, err := services.LikeTweet(body.TweetId, body.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": setLiked,
	})

}

func UnlikeTweet(c *gin.Context) {
	var body structs.TweetUser
	c.Bind(&body)

	setUnLiked, err := services.UnLikeTweet(body.TweetId, body.UserId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Like": setUnLiked,
	})

}

func GetLikesOnTweet(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	totalLikes, err := services.GetLikesOnTweet(params.TweetId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": totalLikes,
	})

}

func SubmitComment(c *gin.Context) {
	var body structs.TweetComment
	c.Bind(&body)

	setComment, err := services.SubmitComment(body.TweetId, body.UserId, body.Content)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Comment": setComment,
	})

}

func ShowCommentsOnTweet(c *gin.Context) {
	var params structs.CommentLimit
	c.Bind(&params)

	allComments, err := services.ShowCommentsOnTweet(params.TweetId, params.Limit)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Comments": allComments,
	})

}

func GetTotalCommentOnTweet(c *gin.Context) {
	var params structs.TweetUser
	c.Bind(&params)

	totalComments, err := services.GetCommentsOnTweet(params.TweetId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": totalComments,
	})

}

func UpdateTweetContent(c *gin.Context) {
	var params models.Tweet
	c.Bind(&params)

	updated, err := services.UpdateContent(params.Id, params.Content)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": updated,
	})

}

func DeleteTweet(c *gin.Context) {
	var body models.Tweet
	c.Bind(&body)

	deleted, err := services.DeleteTweet(body.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Deleted": deleted,
	})

}
