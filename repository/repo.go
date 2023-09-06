// repository.go
package repository

import (
	"fmt"
	"twitter-back-end/initializers"
	"twitter-back-end/models"
	"twitter-back-end/structs"
)

func CreateUser(Id int, Email string, Password string, ThirdParty bool, D_o_b string, FirstName string, LastName string) (bool, error) {

	if D_o_b == "" {
		err := initializers.DB.Exec("INSERT INTO users (first_name,last_name,d_o_b,email,password,third_party) VALUES (?, ?,?,?,?,?)", FirstName, LastName, nil, Email, Password, ThirdParty).Error
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		err := initializers.DB.Exec("INSERT INTO users (first_name,last_name,d_o_b,email,password,third_party) VALUES (?,?,?,?,?,?)", FirstName, LastName, D_o_b, Email, Password, ThirdParty).Error
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func GetUser(Email string) (*models.User, error) {

	var user models.User
	err := initializers.DB.Where("email = ?", Email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidateUser(Email string, password string) (*models.User, error) {

	var user models.User
	err := initializers.DB.Where("email = ? ", Email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetFollowersIds(Id int) ([]int, error) {

	var peopleList []int
	err := initializers.DB.Raw("SELECT follower_id FROM user_followers WHERE user_id = ? LIMIT ?", Id, 5).Scan(&peopleList).Error

	if err != nil {
		return nil, err
	}
	return peopleList, nil
}

func GetPeopleNotFollowed(peopleList []int) ([]models.User, error) {

	var users []models.User
	err := initializers.DB.Not("id", peopleList).Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindOtherUsers(Id int) ([]models.User, error) {

	var users []models.User
	err := initializers.DB.Where("id != ?", Id).Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func AddFollower(UserId int, FollowerId int) (string, error) {

	err := initializers.DB.Exec("INSERT INTO user_followers ( user_id, follower_id) VALUES (?, ?)", UserId, FollowerId).Error

	if err != nil {
		return "", err
	}
	return "Follower Added", nil

}

func GetFollowingsIdsPerPage(Id int, itemsPerPage int, startIndex int) ([]int, error) {

	var peopleList []int
	query := fmt.Sprintf("SELECT follower_id FROM user_followers WHERE user_id = ? LIMIT %d OFFSET %d", itemsPerPage, startIndex)
	err := initializers.DB.Raw(query, Id).Limit(itemsPerPage).Offset(startIndex).Scan(&peopleList).Error

	if err != nil {
		return nil, err
	}
	return peopleList, nil

}

func GetPeoplebyIds(peopleList []int) ([]models.User, error) {

	var users []models.User
	err := initializers.DB.Where("id", peopleList).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil

}

func DeleteFollower(UserId int, FollowerId int) (string, error) {

	err := initializers.DB.Exec("DELETE FROM user_followers where user_id= ? and follower_id = ?", UserId, FollowerId).Error
	if err != nil {
		return "", err
	}
	return "Removed from Following", nil

}

func GetFollowersIdsPerPage(Id int, itemsPerPage int, startIndex int) ([]int, error) {

	var peopleList []int
	query := fmt.Sprintf("SELECT user_id FROM user_followers WHERE follower_id = ? LIMIT %d OFFSET %d", itemsPerPage, startIndex)
	err := initializers.DB.Raw(query, Id).Limit(itemsPerPage).Offset(startIndex).Scan(&peopleList).Error

	if err != nil {
		return nil, err
	}
	return peopleList, nil

}

func UpdateUserInfo(Id int, FirstName string, LastName string, D_o_b string) (bool, error) {

	err := initializers.DB.Exec("UPDATE users SET first_name = ?, last_name = ?, d_o_b = ? WHERE id = ?", FirstName, LastName, D_o_b, Id).Error

	if err != nil {
		return false, err
	}
	return true, nil

}

func SetProfile(Id int, Link string) (bool, error) {

	err := initializers.DB.Exec("UPDATE users SET profile = ? WHERE id = ?", Link, Id).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func GetTotalFollowers(Id int) (int, error) {

	var count int64
	err := initializers.DB.Raw("SELECT COUNT(*) FROM user_followers WHERE follower_id = ?", Id).Scan(&count).Error

	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func GetTotalFollowing(Id int) (int, error) {

	var count int64
	err := initializers.DB.Raw("SELECT COUNT(*) FROM user_followers WHERE user_id = ?", Id).Scan(&count).Error

	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func GetUserById(Id int) (*models.User, error) {

	var user models.User
	err := initializers.DB.Where("Id = ?", Id).Find(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdatePassword(hash string, Id int) (bool, error) {

	err := initializers.DB.Exec("UPDATE users SET password = ? WHERE id = ?", hash, Id).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

//Tweets

func PostTweet(Content string, Id int, Link string) (*models.Tweet, error) {

	tweet := models.Tweet{Content: Content, UserId: Id, Link: Link, Id: 0}
	err := initializers.DB.Create(&tweet).Error

	if err != nil {
		return nil, err
	}
	return &tweet, nil
}

func GetTweetsPerPage(Id int, itemsPerPage int, startIndex int) ([]models.Tweet, error) {

	var tweets []models.Tweet
	err := initializers.DB.Where("user_Id = ?", Id).Order("tweets.id desc").Limit(itemsPerPage).Offset(startIndex).Find(&tweets).Error

	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func TweetLiked(TweetId int, UserId int) (*models.TweetLikes, error) {

	var likeData models.TweetLikes
	err := initializers.DB.Raw("SELECT * FROM tweets_likes WHERE tweet_id = ? AND user_id = ?", TweetId, UserId).Scan(&likeData).Error

	if err != nil {
		return nil, err
	}
	return &likeData, nil
}

func LikeTweet(TweetId int, UserId int) (bool, error) {
	err := initializers.DB.Exec("INSERT INTO tweets_likes ( tweet_id, user_id) VALUES (?, ?)", TweetId, UserId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UnLikeTweet(TweetId int, UserId int) (bool, error) {
	err := initializers.DB.Exec("DELETE FROM tweets_likes WHERE tweet_id = ? AND user_id = ?", TweetId, UserId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetLikesOnTweet(TweetId int) (int, error) {
	var count int64
	err := initializers.DB.Raw("SELECT COUNT(*) FROM tweets_likes WHERE tweet_id = ?", TweetId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
func SubmitComment(TweetId int, UserId int, Content string) (bool, error) {
	err := initializers.DB.Exec("INSERT INTO tweets_comments ( tweet_id, user_id,tweet_comment) VALUES (?, ?,?)", TweetId, UserId, Content).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func ShowComments(TweetId int, Limit int) ([]structs.CommentUser, error) {
	var allComments []structs.CommentUser
	err := initializers.DB.Table("tweets_comments").
		Select("tweets_comments.id, tweets_comments.tweet_id, tweets_comments.user_id, tweets_comments.tweet_comment, users.email, users.first_name, users.last_name , users.profile").
		Joins("left join users on tweets_comments.user_id = users.id").
		Where("tweets_comments.tweet_id = ?", TweetId).
		Order("tweets_comments.id desc").
		Limit(Limit).
		Find(&allComments).Error

	if err != nil {
		return nil, err
	}
	return allComments, nil
}

func GetCommentsOnTweet(TweetId int) (int, error) {
	var count int64
	err := initializers.DB.Raw("SELECT COUNT(*) FROM tweets_comments WHERE tweet_id = ?", TweetId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func UpdateContent(TweetId int, Content string) (bool, error) {
	err := initializers.DB.Exec("UPDATE tweets SET content = ? WHERE id = ?", Content, TweetId).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
func DeleteTweet(TweetId int) (bool, error) {
	err := initializers.DB.Exec("DELETE FROM tweets WHERE id = ?", TweetId).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func GetFollowersTweet(Id int, itemsPerPage int, startIndex int) ([]structs.TweetData, error) {
	var tweets []structs.TweetData
	err := initializers.DB.Table("tweets").
		Joins("LEFT JOIN user_followers ON user_followers.follower_id = tweets.user_id").
		Joins("LEFT JOIN users ON users.id = tweets.user_id").
		Where("(user_followers.user_id = ?) OR (tweets.user_id = ?)", Id, Id).
		Select("tweets.*, users.first_name, users.last_name, users.email, users.profile").
		Order("tweets.id desc").
		Limit(itemsPerPage).
		Offset(startIndex).
		Find(&tweets).Error

	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func SendMessage(sender_id int, reciever_id int, type_of_message string, status string, content string) (bool, error) {
	err := initializers.DB.Exec("INSERT INTO messages (sender_id, reciever_id,message_type,status,content) VALUES (?, ?,?,?,?)", sender_id, reciever_id, type_of_message, status, content).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func GetMessages(sender_id int, reciever_id int) ([]models.Message, error) {
	var messages []models.Message
	err := initializers.DB.Raw("SELECT * FROM messages WHERE sender_id = ? AND reciever_id = ? UNION SELECT * FROM messages WHERE sender_id = ? AND reciever_id = ?", sender_id, reciever_id, reciever_id, sender_id).Scan(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetConversations(userId int) ([]structs.ConversationData, error) {
	var conversations []structs.ConversationData

	rawSQL := `
    (SELECT 
        users.Id as user_id, users.email as user_email, users.first_name as user_first_name, 
        users.last_name as user_last_name, users.profile as user_profile, conversations.* 
    FROM conversations
    INNER JOIN users ON conversations.participant2 = users.id
    WHERE conversations.participant1 = ?)
    
    UNION
    
    (SELECT 
        users.Id as user_id, users.email as user_email, users.first_name as user_first_name, 
        users.last_name as user_last_name, users.profile as user_profile, conversations.* 
    FROM conversations
    INNER JOIN users ON conversations.participant1 = users.id
    WHERE conversations.participant2 = ?)
    
    ORDER BY last_chat DESC
`

	err := initializers.DB.Raw(rawSQL, userId, userId).Scan(&conversations).Error

	if err != nil {
		return nil, err
	}
	return conversations, nil
}
