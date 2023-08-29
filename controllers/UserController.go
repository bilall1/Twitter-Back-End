package controllers

import (
	"net/http"

	"twitter-back-end/models"
	userService "twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
	// PostgreSQL driver
)

func CreateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	userResponse, err := userService.CreateUser(body)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": userResponse,
	})

}

func GetUser(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	userResponse, err := userService.GetUser(params.Email)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": userResponse,
	})

}

func ValidateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	userResponse, err := userService.ValidateUser(body)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": userResponse,
	})

}

func GetPeopleToFollow(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	people, err := userService.GetPeopleToFollow(params.Id)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"people": people,
	})

}

func AddtofollowerList(c *gin.Context) {
	var body structs.UserFollower
	c.Bind(&body)

	followerAdded, err := userService.AddtofollowerList(body.UserId, body.FollowerId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Follower": followerAdded,
	})
}

func GetFollowing(c *gin.Context) {
	var params structs.FollowingPeople
	c.Bind(&params)

	users, err := userService.GetFollowings(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Following": users,
	})

}

func DeleteFollower(c *gin.Context) {
	var body structs.UserFollower
	c.Bind(&body)

	deleted, err := userService.DeleteFollower(body.UserId, body.FollowerId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Message": deleted,
	})

}

func GetFollowers(c *gin.Context) {
	var params structs.FollowerPeople
	c.Bind(&params)

	users, err := userService.GetFollowers(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Followers": users,
	})

}

func UpdateUserData(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	isUpdated, err := userService.UpdateUserData(body)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": isUpdated,
	})

}

func UpdateUserPassword(c *gin.Context) {
	var body structs.Password
	c.Bind(&body)

	isUpdated, err := userService.UpdateUserPassword(body.Id, body.OldPassword, body.NewPassword)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": isUpdated,
	})
}

func AddProfilePicture(c *gin.Context) {
	var body structs.UserProfile
	c.Bind(&body)

	isAdded, err := userService.AddProfilePicture(body.Id, body.Link)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Picture": isAdded,
	})
}

func GetTotalFollowers(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	followersCount, err := userService.GetTotalFollowers(body.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": followersCount,
	})

}
func GetTotalFollowings(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	followingCount, err := userService.GetTotalFollowings(body.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": followingCount,
	})
}

func GenerateToken(c *gin.Context) {

	var body models.User
	c.Bind(&body)

	token, err := userService.GenerateToken(body.Email)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
