// service.go
package services

import (
	"errors"
	"time"
	"twitter-back-end/models"
	"twitter-back-end/repository"
	"twitter-back-end/sharedFunctions"

	"github.com/dgrijalva/jwt-go"
)

func CreateUser(user models.User) (bool, error) {
	hash, _ := sharedFunctions.HashPassword(user.Password)
	userResponse, err := repository.CreateUser(user.Id, user.Email, hash, user.ThirdParty, user.D_o_b, user.FirstName, user.LastName)
	return userResponse, err
}

func GetUser(Email string) (*models.User, error) {

	userResponse, err := repository.GetUser(Email)
	return userResponse, err
}

func ValidateUser(user models.User) (*models.User, error) {

	userResponse, err := repository.ValidateUser(user.Email, user.Password)
	if sharedFunctions.CheckPasswordHash(user.Password, userResponse.Password) {
		return userResponse, err
	} else {
		return nil, errors.New("Password not Matched")
	}
}

func FindOtherUsers(Id int) ([]models.User, error) {
	users, err := repository.FindOtherUsers(Id)
	return users, err
}

func GetPeopleToFollow(Id int) ([]models.User, error) {

	peopleList, err := repository.GetFollowersIds(Id)
	if err != nil {
		return nil, errors.New("Error getting followers Id's list")
	}
	if len(peopleList) > 0 {
		peopleList = append(peopleList, Id)
		peopleToFollow, err := repository.GetPeopleNotFollowed(peopleList)
		return peopleToFollow, err
	} else {
		peopleToFollow, err := repository.FindOtherUsers(Id)
		return peopleToFollow, err
	}

}

func AddtofollowerList(UserId int, FollowerId int) (string, error) {

	followerAdded, err := repository.AddFollower(UserId, FollowerId)
	return followerAdded, err
}

func GetFollowings(Id int, Page int) ([]models.User, error) {

	itemsPerPage := 3
	startIndex := (Page - 1) * itemsPerPage

	peopleList, err := repository.GetFollowingsIdsPerPage(Id, itemsPerPage, startIndex)
	if err != nil {
		return nil, errors.New("Error getting following Id's per page list")
	}
	if len(peopleList) == 0 {
		return []models.User{}, nil
	}
	users, err := repository.GetPeoplebyIds(peopleList)
	return users, err

}

func DeleteFollower(UserId int, FollowerId int) (string, error) {

	deleted, err := repository.DeleteFollower(UserId, FollowerId)
	return deleted, err
}

func GetFollowers(Id int, Page int) ([]models.User, error) {

	itemsPerPage := 3
	startIndex := (Page - 1) * itemsPerPage

	peopleList, err := repository.GetFollowersIdsPerPage(Id, itemsPerPage, startIndex)

	if err != nil {
		return nil, errors.New("Error getting followers Id's per page list")
	}
	if len(peopleList) == 0 {
		return []models.User{}, nil
	}
	users, err := repository.GetPeoplebyIds(peopleList)
	return users, err
}

func UpdateUserData(user models.User) (bool, error) {
	isUpdated, err := repository.UpdateUserInfo(user.Id, user.FirstName, user.LastName, user.D_o_b)
	return isUpdated, err
}

func UpdateUserPassword(Id int, OldPassword string, NewPassword string) (bool, error) {
	user, err := repository.GetUserById(Id)
	if err != nil {
		return false, errors.New("Error getting user ")
	}

	if sharedFunctions.CheckPasswordHash(OldPassword, user.Password) {
		newHash, _ := sharedFunctions.HashPassword(NewPassword)
		isSet, err := repository.UpdatePassword(newHash, Id)
		return isSet, err
	}
	return false, nil
}

func AddProfilePicture(Id int, Link string) (bool, error) {
	isAdded, err := repository.SetProfile(Id, Link)
	return isAdded, err
}

func GetTotalFollowers(Id int) (int, error) {
	followersCount, err := repository.GetTotalFollowers(Id)
	return followersCount, err
}

func GetTotalFollowings(Id int) (int, error) {
	followingCount, err := repository.GetTotalFollowing(Id)
	return followingCount, err
}

func GenerateToken(Email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": Email,                                 // User identifier
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	tokenString, err := token.SignedString([]byte("aurorasolutions"))
	if err != nil {
		return tokenString, errors.New("Internal Server Error")
	}
	return tokenString, nil
}

func GetStatus(user_id int) (models.UserStatus, error) {
	status, err := repository.GetStatus(user_id)
	return status, err
}

func GetOnlineStatus(user_id int) ([]models.UserStatus, error) {
	allUserStatus, err := repository.GetOnlineStatus(user_id)
	return allUserStatus, err
}

func UpdateStatus(user_id int, status string) (bool, error) {
	isSet, err := repository.UpdateStatus(user_id, status)
	return isSet, err
}

func UpdateNotificationToken(user_id int, token string) (bool, error) {
	isSet, err := repository.UpdateNotificationToken(user_id, token)
	return isSet, err
}

func GetUserNotification(user_id int) (*models.UserNotification, error) {
	user, err := repository.GetUserNotification(user_id)
	return user, err
}

func GetUserById(user_id int) (*models.User, error) {
	user, err := repository.GetUserById(user_id)
	return user, err
}
