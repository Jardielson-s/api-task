package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
)

// FindUserById godoc
//
//	@Summary		Find user by id
//	@Description	Find user by id in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "User ID"
// @Success 200 {object} shared.CreateUserResponse "User data"
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "User not found"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/users/{id} [get]
func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)

	result, err := h.repo.FindById(idNum)

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
	_ = json.NewEncoder(w).Encode(result)
}
