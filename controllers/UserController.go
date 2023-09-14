package controllers

import (
	"net/http"

	"twitter-back-end/models"
	services "twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
	// PostgreSQL driver
)

func CreateUser(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	userResponse, err := services.CreateUser(body)

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

	userResponse, err := services.GetUser(params.Email)

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

	userResponse, err := services.ValidateUser(body)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": userResponse,
	})

}

func FindOtherUsers(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	userResponse, err := services.FindOtherUsers(params.Id)
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

	peopleResponse, err := services.GetPeopleToFollow(params.Id)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"people": peopleResponse,
	})

}

func AddtofollowerList(c *gin.Context) {
	var body structs.UserFollower
	c.Bind(&body)

	followerAddedResponse, err := services.AddtofollowerList(body.UserId, body.FollowerId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Follower": followerAddedResponse,
	})
}

func GetFollowing(c *gin.Context) {
	var params structs.FollowingPeople
	c.Bind(&params)

	usersResponse, err := services.GetFollowings(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Following": usersResponse,
	})

}

func DeleteFollower(c *gin.Context) {
	var body structs.UserFollower
	c.Bind(&body)

	deletedResponse, err := services.DeleteFollower(body.UserId, body.FollowerId)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Message": deletedResponse,
	})

}

func GetFollowers(c *gin.Context) {
	var params structs.FollowerPeople
	c.Bind(&params)

	usersResponse, err := services.GetFollowers(params.Id, params.Page)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Followers": usersResponse,
	})

}

func UpdateUserData(c *gin.Context) {
	var body models.User
	c.Bind(&body)

	isUpdatedResponse, err := services.UpdateUserData(body)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": isUpdatedResponse,
	})

}

func UpdateUserPassword(c *gin.Context) {
	var body structs.Password
	c.Bind(&body)

	isUpdatedResponse, err := services.UpdateUserPassword(body.Id, body.OldPassword, body.NewPassword)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"update": isUpdatedResponse,
	})
}

func AddProfilePicture(c *gin.Context) {
	var body structs.UserProfile
	c.Bind(&body)

	isAddedResponse, err := services.AddProfilePicture(body.Id, body.Link)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Picture": isAddedResponse,
	})
}

func GetTotalFollowers(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	followersCountResponse, err := services.GetTotalFollowers(params.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": followersCountResponse,
	})

}
func GetTotalFollowings(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	followingCountResponse, err := services.GetTotalFollowings(params.Id)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Count": followingCountResponse,
	})
}

func GenerateToken(c *gin.Context) {

	var body models.User
	c.Bind(&body)

	token, err := services.GenerateToken(body.Email)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func GetStatus(c *gin.Context) {

	var body models.UserStatus
	c.Bind(&body)

	status, err := services.GetStatus(body.UserId)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status": status,
	})

}

func GetOnlineStatus(c *gin.Context) {

	var body models.User
	c.Bind(&body)

	allUserStatus, err := services.GetOnlineStatus(body.Id)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status": allUserStatus,
	})

}

func UpdateStatus(c *gin.Context) {

	var params models.UserStatus
	c.Bind(&params)

	setStatusResponse, err := services.UpdateStatus(params.UserId, params.Status)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status": setStatusResponse,
	})

}
