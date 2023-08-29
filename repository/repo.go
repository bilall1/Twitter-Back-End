// repository.go
package repository

import (
	"fmt"
	"twitter-back-end/initializers"
	"twitter-back-end/models"
)

func CreateUser(Id int, Email string, Password string, ThirdParty bool, D_o_b string, FirstName string, LastName string) (*models.User, error) {

	user := models.User{FirstName: FirstName, LastName: LastName, D_o_b: D_o_b, Email: Email, Password: Password, ThirdParty: ThirdParty, Id: 0}
	err := initializers.DB.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(Email string) (*models.User, error) {

	var user models.User
	err := initializers.DB.Debug().Where("email = ?", Email).Find(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidateUser(Email string, password string) (*models.User, error) {

	var user models.User
	err := initializers.DB.Debug().Where("email = ? ", Email).First(&user).Error

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

	err := initializers.DB.Debug().Exec("UPDATE users SET first_name = ?, last_name = ?, d_o_b = ? WHERE id = ?", FirstName, LastName, D_o_b, Id).Error

	if err != nil {
		return false, err
	}
	return true, nil

}

func SetProfile(Id int, Link string) (bool, error) {

	err := initializers.DB.Debug().Exec("UPDATE users SET profile = ? WHERE id = ?", Link, Id).Error

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
	err := initializers.DB.Debug().Where("Id = ?", Id).Find(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdatePassword(hash string, Id int) (bool, error) {

	err := initializers.DB.Debug().Exec("UPDATE users SET password = ? WHERE id = ?", hash, Id).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
