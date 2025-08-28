package handlers

import (
	"encoding/json"
	"net/http"

	entity "github.com/Jardielson-s/api-task/modules/auth/entities"
	"github.com/Jardielson-s/api-task/modules/auth/services"
	"github.com/Jardielson-s/api-task/modules/users/repository"
)

type LoginHandler struct {
	service  services.AuthService
	userRepo repository.UserRepository
}

// Login godoc
//
//	@Summary		Login
//	@Description	route to login
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Param			login	body		entity.Login	true	"Login data"
//	@Success		201		{object}	entity.Login
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/auth/login [post]
func NewLoginHandler(s services.AuthService, userR repository.UserRepository) *LoginHandler {
	return &LoginHandler{s, userR}
}

func (h *LoginHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login entity.Login
	if err := json.NewDecoder((r.Body)).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
