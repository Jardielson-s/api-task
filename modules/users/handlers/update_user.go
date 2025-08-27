package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
	"github.com/Jardielson-s/api-task/modules/users/model"
)

type UpdateUserBody struct {
	Email    string  `json:"email" validate:"required,min=10,max=100"`
	Password *string `json:"password" validate:"required,min=6,max=8"`
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update user in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "User ID"
//
//	@Param			user	body		UpdateUserBody true "Descrição do parâmetro"
//	@Success		201		{object}	shared.CreateUserResponse  "Usuário atualizado com sucesso"
//	@Failure		400		{object}	shared.CreateResponse
//
// @Failure 		400 	{object}	shared.CreateResponse
//
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/users/update/{id} [patch]
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UpdateUserBody
	id := r.URL.Path[len("/users/update/"):]
	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)
	if err := json.NewDecoder((r.Body)).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(user)
	fmt.Println(err)
	if err != nil {
		errorResponse := shared.CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: err.Error(),
		}
		shared.JSONError(w, errorResponse, http.StatusBadRequest)
		return
	}
	result, err := h.service.UpdateUserService(idNum, model.UpdateUser{
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
