package handlers

import (
	"encoding/json"
	"net/http"
	"subscription-management/models"
	"subscription-management/repository"
	"subscription-management/utils"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AdminHandler struct {
	rdb    *redis.Client
	logger *zap.Logger
}

func NewAdminHandler(rdb *redis.Client, logger *zap.Logger) *AdminHandler {
	return &AdminHandler{rdb: rdb, logger: logger}
}

func (h *AdminHandler) CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.CreateUser(user)
	h.logger.Info("User created", zap.String("username", user.Username))

	userData, err := json.Marshal(user)
	if err != nil {
		h.logger.Error("Error marshalling user data", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, "Error marshalling user data")
		return
	}

	err = h.rdb.Set(utils.GetContext(), user.Username, userData, 0).Err()
	if err != nil {
		h.logger.Error("Error saving data to Redis", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, "Error saving data to Redis")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *AdminHandler) UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err))
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.UpdateUser(user)
	h.logger.Info("User updated", zap.String("username", user.Username))

	userData, err := json.Marshal(user)
	if err != nil {
		h.logger.Error("Error marshalling user data", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, "Error marshalling user data")
		return
	}

	err = h.rdb.Set(utils.GetContext(), user.Username, userData, 0).Err()
	if err != nil {
		h.logger.Error("Error updating data in Redis", zap.Error(err))
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating data in Redis")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
