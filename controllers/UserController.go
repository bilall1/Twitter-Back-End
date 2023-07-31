package controllers

import (
	"fmt"

	"github.com/bilall1/twitter-backend/initializers"
	"github.com/bilall1/twitter-backend/models"
	"github.com/gin-gonic/gin"
	// PostgreSQL driver
)

func CreateUser(c *gin.Context) {

	var body struct {
		Id         int
		Email      string
		Password   string
		ThirdParty bool
		D_o_b      string
		FirstName  string
		LastName   string
	}
	c.Bind(&body)

	user := models.User{FirstName: body.FirstName, LastName: body.LastName, D_o_b: body.D_o_b, Email: body.Email, Password: body.Password, ThirdParty: body.ThirdParty, Id: 0}

	result := initializers.DB.Debug().Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Normal response object
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUser(c *gin.Context) {

	var body struct {
		Email    string
	}
	c.Bind(&body)

	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})

}

func ValidateUser(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	var user models.User
	userobject := initializers.DB.Debug().Where("email = ? AND password = ?", body.Email, body.Password).First(&user)

	if userobject.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})

}

func GetPeopleToFollow(c *gin.Context) {

	var body struct {
		Email string
	}
	c.Bind(&body)

	//Fetching User Id from Email
	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Getting id's of people that appear in people you may know
	var user_followers []int
	result1 := initializers.DB.Raw("SELECT follower_id FROM user_followers WHERE user_id = ?", user.Id).Scan(&user_followers)

	if result1.Error != nil {
		c.Status(400)
		return
	}

	//Getting specific users array from that id's

	var users []models.User
	if len(user_followers) > 0 {

		user_followers = append(user_followers, user.Id)

		result2 := initializers.DB.Not("id", user_followers).Find(&users)
		if result2.Error != nil {
			c.Status(400)
			return
		}
	} else {
		result2 := initializers.DB.Where("id != ?", user.Id).Find(&users)
		if result2.Error != nil {
			c.Status(400)
			return
		}

	}
	if len(users) <= 5 {

	} else {
		users = users[0:5]

	}

	c.JSON(200, gin.H{
		"people": users,
	})

}

func AddtofollowerList(c *gin.Context) {

	var body struct {
		Id         int
		Email      string
		FollowerID int
	}
	c.Bind(&body)

	//Fetching User Id from Email
	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result1 := initializers.DB.Exec("INSERT INTO user_followers ( user_id, follower_id) VALUES (?, ?)", user.Id, body.FollowerID)

	if result1.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Follower": "Follower Added",
	})
}

func GetFollowing(c *gin.Context) {

	var body struct {
		Email string
		Page  int
	}
	c.Bind(&body)

	itemsPerPage := 3

	startIndex := (body.Page - 1) * itemsPerPage

	//Fetching User Id from Email
	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Getting id's of people that appear in people you may know
	var user_followers []int
	query := fmt.Sprintf("SELECT follower_id FROM user_followers WHERE user_id = ? LIMIT %d OFFSET %d", itemsPerPage, startIndex)
	result1 := initializers.DB.Raw(query, user.Id).Limit(itemsPerPage).Offset(startIndex).Scan(&user_followers)

	if result1.Error != nil {
		c.Status(400)
		return
	}

	if len(user_followers) == 0 {

		c.JSON(200, gin.H{
			"Following": []models.User{},
		})

	}

	var users []models.User
	result2 := initializers.DB.Where("id", user_followers).Find(&users)
	if result2.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Following": users,
	})

}

func DeleteFollower(c *gin.Context) {

	var body struct {
		Email      string
		FollowerId int
	}
	c.Bind(&body)

	//Fetching User Id from Email
	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	fmt.Println(user.Id, body.FollowerId)
	initializers.DB.Exec("DELETE FROM user_followers where user_id= ? and follower_id = ?", user.Id, body.FollowerId)

	c.JSON(200, gin.H{
		"Message": "Removed from Following",
	})

}

func GetFollowers(c *gin.Context) {

	var body struct {
		Email string
		Page  int
	}
	c.Bind(&body)

	itemsPerPage := 3

	startIndex := (body.Page - 1) * itemsPerPage

	//Fetching User Id from Email
	var user models.User
	result := initializers.DB.Debug().Where("email = ?", body.Email).Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Getting id's of people that appear in people you may know
	var user_followers []int
	query := fmt.Sprintf("SELECT user_id FROM user_followers WHERE follower_id = ? LIMIT %d OFFSET %d", itemsPerPage, startIndex)
	result1 := initializers.DB.Raw(query, user.Id).Limit(itemsPerPage).Offset(startIndex).Scan(&user_followers)

	if result1.Error != nil {
		c.Status(400)
		return
	}
	if len(user_followers) == 0 {

		c.JSON(200, gin.H{
			"Followers": []models.User{},
		})

	}

	var users []models.User
	result2 := initializers.DB.Where("id", user_followers).Find(&users)
	if result2.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Followers": users,
	})

}
