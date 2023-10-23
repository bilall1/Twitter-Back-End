package controllers

import (
	"net/http"

	"twitter-back-end/models"
	services "twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
	// PostgreSQL driver
)

func CreateUser(ctx *gin.Context) {
	var body models.User
	ctx.Bind(&body)

	userResponse, err := services.CreateUser(body)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": userResponse,
	})

}

func GetUser(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	userResponse, err := services.GetUser(params.Email)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": userResponse,
	})

}

func ValidateUser(ctx *gin.Context) {
	var body models.User
	ctx.Bind(&body)

	userResponse, err := services.ValidateUser(body)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": userResponse,
	})

}

func FindOtherUsers(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	userResponse, err := services.FindOtherUsers(params.Id)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": userResponse,
	})

}

func GetPeopleToFollow(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	peopleResponse, err := services.GetPeopleToFollow(params.Id)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"people": peopleResponse,
	})

}

func AddtofollowerList(ctx *gin.Context) {
	var body structs.UserFollower
	ctx.Bind(&body)

	followerAddedResponse, err := services.AddtofollowerList(body.UserId, body.FollowerId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Follower": followerAddedResponse,
	})
}

func GetFollowing(ctx *gin.Context) {
	var params structs.FollowingPeople
	ctx.Bind(&params)

	usersResponse, err := services.GetFollowings(params.Id, params.Page)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Following": usersResponse,
	})

}

func DeleteFollower(ctx *gin.Context) {
	var body structs.UserFollower
	ctx.Bind(&body)

	deletedResponse, err := services.DeleteFollower(body.UserId, body.FollowerId)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Message": deletedResponse,
	})

}

func GetFollowers(ctx *gin.Context) {
	var params structs.FollowerPeople
	ctx.Bind(&params)

	usersResponse, err := services.GetFollowers(params.Id, params.Page)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Followers": usersResponse,
	})

}

func UpdateUserData(ctx *gin.Context) {
	var body models.User
	ctx.Bind(&body)

	isUpdatedResponse, err := services.UpdateUserData(body)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"update": isUpdatedResponse,
	})

}

func UpdateUserPassword(ctx *gin.Context) {
	var body structs.Password
	ctx.Bind(&body)

	isUpdatedResponse, err := services.UpdateUserPassword(body.Id, body.OldPassword, body.NewPassword)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"update": isUpdatedResponse,
	})
}

func AddProfilePicture(ctx *gin.Context) {
	var body structs.UserProfile
	ctx.Bind(&body)

	isAddedResponse, err := services.AddProfilePicture(body.Id, body.Link)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Picture": isAddedResponse,
	})
}

func GetTotalFollowers(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	followersCountResponse, err := services.GetTotalFollowers(params.Id)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Count": followersCountResponse,
	})

}
func GetTotalFollowings(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	followingCountResponse, err := services.GetTotalFollowings(params.Id)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"Count": followingCountResponse,
	})
}

func GenerateToken(ctx *gin.Context) {

	var body models.User
	ctx.Bind(&body)

	token, err := services.GenerateToken(body.Email)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func GetStatus(ctx *gin.Context) {

	var body models.UserStatus
	ctx.Bind(&body)

	status, err := services.GetStatus(body.UserId)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"status": status,
	})

}

func GetOnlineStatus(ctx *gin.Context) {

	var body models.User
	ctx.Bind(&body)

	allUserStatus, err := services.GetOnlineStatus(body.Id)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"status": allUserStatus,
	})

}

func UpdateStatus(ctx *gin.Context) {

	var params models.UserStatus
	ctx.Bind(&params)

	setStatusResponse, err := services.UpdateStatus(params.UserId, params.Status)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"status": setStatusResponse,
	})

}

func UpdateNotificationToken(ctx *gin.Context) {

	var params models.UserNotification
	ctx.Bind(&params)

	setStatusResponse, err := services.UpdateNotificationToken(params.UserId, params.Token)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"status": setStatusResponse,
	})

}
