package handlers

import (
	"encoding/json"
	"net/http"

	entity "github.com/Jardielson-s/api-task/modules/users/entities"
	"github.com/Jardielson-s/api-task/modules/users/services"
)

type UserHandler struct {
	service services.UserService
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entity.User	true	"User data"
//	@Success		201		{object}	entity.User
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/users [post]
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder((r.Body)).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.service.CreateUserService(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
