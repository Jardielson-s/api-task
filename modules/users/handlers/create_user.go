package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
	userRolesModel "github.com/Jardielson-s/api-task/modules/user_roles/repository"
	"github.com/Jardielson-s/api-task/modules/users/model"
	"github.com/Jardielson-s/api-task/modules/users/repository"
	"github.com/Jardielson-s/api-task/modules/users/services"
	"github.com/go-playground/validator/v10"
)

type CreateUserBody struct {
	Username string `json:"username" validate:"required,min=5,max=100" example:"Test User"`
	Email    string `json:"email" validate:"required,min=10,max=100,email" example:"test@gmail.com"`
	Password string `json:"password" validate:"required,min=6,max=8" example:"test1234"`
}

type UserHandler struct {
	service       services.UserService
	repo          repository.UserRepository
	userRolesRepo userRolesModel.UserRolesRepository
}

var validate = validator.New()

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		CreateUserBody true "Descrição do parâmetro"
//	@Success		201		{object}	shared.CreateUserResponse  "Usuário criado com sucesso"
//	@Failure		400		{object}	shared.CreateResponse
//
// @Failure 		400 	{object}	shared.CreateResponse
//
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/users [post]
func NewUserHandler(service services.UserService, repository repository.UserRepository, userRolesRepo userRolesModel.UserRolesRepository) *UserHandler {
	return &UserHandler{service, repository, userRolesRepo}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user CreateUserBody
	if err := json.NewDecoder((r.Body)).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(user)
	if err != nil {
		errorResponse := shared.CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: err.Error(),
		}
		shared.JSONError(w, errorResponse, http.StatusBadRequest)
		return
	}
	result, err := h.service.CreateUserService(&model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		errorResponse := shared.CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: err.Error(),
		}
		shared.JSONError(w, errorResponse, http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
