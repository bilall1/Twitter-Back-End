package services

import (
	"errors"
	"twitter-back-end/models"
	"twitter-back-end/repository"
	"twitter-back-end/structs"
)

func PostTweet(Content string, Id int, Link string) (*models.Tweet, error) {
	tweet, err := repository.PostTweet(Content, Id, Link)
	return tweet, err
}

func GetTweet(Email string, Page int) ([]models.Tweet, error) {
	itemsPerPage := 10
	startIndex := (Page - 1) * itemsPerPage

	user, err := repository.GetUser(Email)
	if err != nil {
		return nil, errors.New("Error getting User")
	}

	tweets, errTweets := repository.GetTweetsPerPage(user.Id, itemsPerPage, startIndex)
	return tweets, errTweets

}

func GetIfTweetLiked(TweetId int, UserId int) (bool, error) {
	likeData, err := repository.TweetLiked(TweetId, UserId)
	if err != nil {
		return false, errors.New("Error getting Like Data")
	}
	if likeData.Id == 0 {
		return false, nil
	} else {
		return true, nil
	}

}

func LikeTweet(TweetId int, UserId int) (bool, error) {
	setLiked, err := repository.LikeTweet(TweetId, UserId)
	return setLiked, err
}

func UnLikeTweet(TweetId int, UserId int) (bool, error) {
	unLiked, err := repository.UnLikeTweet(TweetId, UserId)
	return unLiked, err
}

func GetLikesOnTweet(TweetId int) (int, error) {
	totalLikes, err := repository.GetLikesOnTweet(TweetId)
	if err != nil {
		return 0, err
	}
	return totalLikes, nil
}

func SubmitComment(TweetId int, UserId int, Content string) (bool, error) {
	setComment, err := repository.SubmitComment(TweetId, UserId, Content)
	return setComment, err
}

func ShowCommentsOnTweet(TweetId int, Limit int) ([]structs.CommentUser, error) {
	allComments, err := repository.ShowComments(TweetId, Limit)
	return allComments, err
}

func GetCommentsOnTweet(TweetId int) (int, error) {
	totalComments, err := repository.GetCommentsOnTweet(TweetId)
	if err != nil {
		return 0, err
	}
	return totalComments, nil
}
func UpdateContent(TweetId int, Content string) (bool, error) {
	updated, err := repository.UpdateContent(TweetId, Content)
	return updated, err
}
func DeleteTweet(TweetId int) (bool, error) {
	deleted, err := repository.DeleteTweet(TweetId)
	return deleted, err
}

func GetFollowersTweet(Id int, Page int) ([]structs.TweetData, error) {
	itemsPerPage := 10
	startIndex := (Page - 1) * itemsPerPage

	tweets, err := repository.GetFollowersTweet(Id, itemsPerPage, startIndex)
	if len(tweets) == 0 {
		return []structs.TweetData{}, err
	}
	return tweets, nil
}
