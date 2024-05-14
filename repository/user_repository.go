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

func UpdateUser(updatedUser models.User) {
	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = updatedUser
			return
		}
	}
}

func GetAllUsers() []models.User {
	return users
}
