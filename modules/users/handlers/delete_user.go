package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
)

// DeleteUser(Soft delete) godoc
//
//	@Summary		Delete user
//	@Description	Delete  user in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "User ID"
// @Success 204
// @Failure 404 {object} map[string]string "User not found"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/users/delete/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/delete/"):]
	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)

	err := h.repo.DeleteUser(idNum)

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
