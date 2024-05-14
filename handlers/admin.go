package handlers

import (
	"encoding/json"
	"net/http"
	"subscription-management/models"
	"subscription-management/repository"
	"subscription-management/utils"
)

// CreateServiceHandler creates a new service and stores user data in Redis
func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.CreateUser(user)

	// Store user data in Redis
	rdb := utils.GetRedisClient()
	userData, err := json.Marshal(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error marshalling user data")
		return
	}

	err = rdb.Set(utils.GetContext(), user.Username, userData, 0).Err()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error saving data to Redis")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// UpdateServiceHandler updates an existing service and stores updated user data in Redis
func UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.UpdateUser(user)

	// Update user data in Redis
	rdb := utils.GetRedisClient()
	userData, err := json.Marshal(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error marshalling user data")
		return
	}

	err = rdb.Set(utils.GetContext(), user.Username, userData, 0).Err()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating data in Redis")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
