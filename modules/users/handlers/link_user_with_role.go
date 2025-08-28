package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
)

type LinkUserWithRoleBody struct {
	RoleId int `json:"role_id" validate:"required,gt=0" example:"1"`
	UserId int `json:"user_id" validate:"required,gt=0" example:"1"`
}

// LinkUserWithRole godoc
//
//	@Summary		Link user with role
//	@Description	link user  with role in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security  Bearer
//
//	@Param			user	body		LinkUserWithRoleBody true "Descrição do parâmetro"
//	@Success		204
//	@Failure		400		{object}	shared.CreateResponse
//
// @Failure 		400 	{object}	shared.CreateResponse
//
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/users/link [patch]
func (h *UserHandler) LinkUserWithRoleHandler(w http.ResponseWriter, r *http.Request) {
	var body LinkUserWithRoleBody
	if err := json.NewDecoder((r.Body)).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(body)
	if err != nil {
		errorResponse := shared.CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: err.Error(),
		}
		shared.JSONError(w, errorResponse, http.StatusBadRequest)
		return
	}
	err = h.userRolesRepo.LinkUserWithRole(body.UserId, body.RoleId)

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
	w.WriteHeader(http.StatusNoContent)
}
