package repository

import (
	"subscription-management/models"
)

var users []models.User

func CreateUser(user models.User) {
	users = append(users, user)
}

func GetUserByUsername(username string) *models.User {
	for _, user := range users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func UpdateUser(user models.User) {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
			return
		}
	}
}

func GetAllUsers() []models.User {
	return users
}
