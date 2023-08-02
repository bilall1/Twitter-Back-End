package controllers

import (
	"fmt"

	"github.com/bilall1/twitter-backend/initializers"
	"github.com/bilall1/twitter-backend/models"
	"github.com/gin-gonic/gin"
	// PostgreSQL driver
)

// Struct Defined to send Tweet data.
type tweetData struct {
	Id        int
	Content   string
	UserId    int
	FirstName string
	LastName  string
	Email     string
}

func PostTweet(c *gin.Context) {

	var body struct {
		Id      int
		Content string
	}
	c.Bind(&body)

	tweet := models.Tweet{Content: body.Content, UserId: body.Id, Id: 0}

	result := initializers.DB.Create(&tweet)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, "Tweet Posted")

}

func GetTweet(c *gin.Context) {

	//Retrieving Id from Email
	var body struct {
		Email string
		Page  int
	}
	c.Bind(&body)

	itemsPerPage := 10

	startIndex := (body.Page - 1) * itemsPerPage

	var user models.User
	userobject := initializers.DB.Where("email = ?", body.Email).First(&user)
	if userobject.Error != nil {
		c.Status(400)
		return
	}

	var tweets []models.Tweet

	result := initializers.DB.Where("user_Id = ?", user.Id).Order("tweets.id desc").Limit(itemsPerPage).Offset(startIndex).Find(&tweets)
	if result.Error != nil {
		c.Status(400)
		return
	}
	if len(tweets) == 0 {

		c.JSON(200, gin.H{
			"Tweets":    []models.Tweet{},
			"FirstName": user.FirstName,
			"LastName":  user.LastName,
			"Email":     user.Email,
		})
		return

	}

	c.JSON(200, gin.H{
		"Tweets":    tweets,
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
		"Email":     user.Email,
	})
}

func GetFollowersTweet(c *gin.Context) {

	itemsPerPage := 10

	var body struct {
		Id   int
		Page int
	}
	c.Bind(&body)

	// Calculate the start and end index
	startIndex := (body.Page - 1) * itemsPerPage

	var tweets []models.Tweet
	result := initializers.DB.Table("tweets").
		Joins("LEFT JOIN user_followers ON user_followers.follower_id = tweets.user_id").
		Where("(user_followers.user_id = ?) OR (tweets.user_id = ?)", body.Id, body.Id).
		Order("tweets.id desc").
		Limit(itemsPerPage).
		Offset(startIndex).
		Find(&tweets)

	if len(tweets) == 0 {

		c.JSON(200, gin.H{
			"Tweets": []tweetData{},
		})
		return

	}

	if result.Error != nil {
		c.Status(400)
		return
	}

	var sendTweet []tweetData

	for i := 0; i < len(tweets); i++ {

		var user1 models.User
		result := initializers.DB.Where("Id = ?", tweets[i].UserId).Find(&user1)

		if result.Error != nil {
			c.Status(400)
			return
		}

		var singleTweet tweetData
		singleTweet.Id = tweets[i].Id
		singleTweet.Content = tweets[i].Content
		singleTweet.UserId = tweets[i].UserId
		singleTweet.FirstName = user1.FirstName
		singleTweet.LastName = user1.LastName
		singleTweet.Email = user1.Email

		sendTweet = append(sendTweet, singleTweet)

	}

	c.JSON(200, gin.H{
		"Tweets": sendTweet,
	})

}

func GetIfTweetLiked(c *gin.Context) {

	var body struct {
		TweetId int
		UserId  int
	}
	c.Bind(&body)

	type likeData struct {
		Id      int
		TweetId int
		UserId  int
	}

	var data likeData

	// userobject := initializers.DB.Where("tweet_id = ? AND user_id= ? ", body.tweetId, body.UserId).Scan(&data)

	fmt.Println(body.TweetId, body.UserId)

	result := initializers.DB.Raw("SELECT * FROM tweets_likes WHERE tweet_id = ? AND user_id = ?", body.TweetId, body.UserId).Scan(&data)

	if result.Error != nil {
		c.Status(400)
		return
	}

	fmt.Println(data)

	if data.Id == 0 {

		c.JSON(200, gin.H{
			"Like": 0,
		})

	} else {
		c.JSON(200, gin.H{
			"Like": 1,
		})

	}

}

func LikeTweet(c *gin.Context) {

	var body struct {
		TweetId int
		UserId  int
	}
	c.Bind(&body)

	result1 := initializers.DB.Exec("INSERT INTO tweets_likes ( tweet_id, user_id) VALUES (?, ?)", body.TweetId, body.UserId)

	if result1.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Like": 1,
	})

}

func UnlikeTweet(c *gin.Context) {

	var body struct {
		TweetId int
		UserId  int
	}
	c.Bind(&body)

	result1 := initializers.DB.Exec("DELETE FROM tweets_likes WHERE tweet_id = ? AND user_id = ?", body.TweetId, body.UserId)

	if result1.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"UnLike": 1,
	})

}

func GetLikesOnTweet(c *gin.Context) {

	var body struct {
		TweetId int
	}
	c.Bind(&body)

	var count int64
	result := initializers.DB.Raw("SELECT COUNT(*) FROM tweets_likes WHERE tweet_id = ?", body.TweetId).Scan(&count)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Count": count,
	})

}
