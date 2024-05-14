package handlers

import (
	"encoding/json"
	"net/http"
	"subscription-management/models"
	"subscription-management/repository"
	"subscription-management/utils"
)

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.CreateUser(user)
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.UpdateUser(user)
	utils.RespondWithJSON(w, http.StatusOK, user)
}
