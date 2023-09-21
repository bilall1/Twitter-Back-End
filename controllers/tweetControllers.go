package controllers

import (
	"twitter-back-end/models"
	"twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
)

func PostTweet(ctx *gin.Context) {
	var body models.Tweet
	ctx.Bind(&body)

	tweetResponse, err := services.PostTweet(body.Content, body.Id, body.Link)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Tweet": tweetResponse,
	})

}

func GetTweet(ctx *gin.Context) {
	var params structs.TweetOfUser
	ctx.Bind(&params)

	tweetsResponse, err := services.GetTweet(params.Email, params.Page)
	if err != nil {
		ctx.Status(400)
		return
	}
	if len(tweetsResponse) == 0 {
		ctx.JSON(200, gin.H{
			"Tweets": []models.Tweet{},
		})
		return
	}
	ctx.JSON(200, gin.H{
		"Tweets": tweetsResponse,
	})
}

func GetFollowersTweet(ctx *gin.Context) {
	var params structs.TweetFollower
	ctx.Bind(&params)

	tweetsResponse, err := services.GetFollowersTweet(params.Id, params.Page)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Tweets": tweetsResponse,
	})

}

func GetIfTweetLiked(ctx *gin.Context) {
	var params structs.TweetUser
	ctx.Bind(&params)

	isLikedResponse, err := services.GetIfTweetLiked(params.TweetId, params.UserId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Like": isLikedResponse,
	})

}

func LikeTweet(ctx *gin.Context) {
	var body structs.TweetUser
	ctx.Bind(&body)

	setLikedResponse, err := services.LikeTweet(body.TweetId, body.UserId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Like": setLikedResponse,
	})

}

func UnlikeTweet(ctx *gin.Context) {
	var body structs.TweetUser
	ctx.Bind(&body)

	setUnLikedResponse, err := services.UnLikeTweet(body.TweetId, body.UserId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Like": setUnLikedResponse,
	})

}

func GetLikesOnTweet(ctx *gin.Context) {
	var params structs.TweetUser
	ctx.Bind(&params)

	totalLikesResponse, err := services.GetLikesOnTweet(params.TweetId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Count": totalLikesResponse,
	})

}

func SubmitComment(ctx *gin.Context) {
	var body structs.TweetComment
	ctx.Bind(&body)

	setCommentResponse, err := services.SubmitComment(body.TweetId, body.UserId, body.Content)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Comment": setCommentResponse,
	})

}

func ShowCommentsOnTweet(ctx *gin.Context) {
	var params structs.CommentLimit
	ctx.Bind(&params)

	allCommentsResponse, err := services.ShowCommentsOnTweet(params.TweetId, params.Limit)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Comments": allCommentsResponse,
	})

}

func GetTotalCommentOnTweet(ctx *gin.Context) {
	var params structs.TweetUser
	ctx.Bind(&params)

	totalCommentsResponse, err := services.GetCommentsOnTweet(params.TweetId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Count": totalCommentsResponse,
	})

}

func UpdateTweetContent(ctx *gin.Context) {
	var params models.Tweet
	ctx.Bind(&params)

	updatedResponse, err := services.UpdateContent(params.Id, params.Content)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"update": updatedResponse,
	})

}

func DeleteTweet(ctx *gin.Context) {
	var body models.Tweet
	ctx.Bind(&body)

	deletedResponse, err := services.DeleteTweet(body.Id)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Deleted": deletedResponse,
	})

}
