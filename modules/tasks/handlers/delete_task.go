package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/modules/shared"
)

// DeleteTask(Soft delete) godoc
//
//	@Summary		Delete task
//	@Description	Delete  task in the database
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
// @Security  Bearer
//
// @Param id path int true "Task ID"
// @Success 204
// @Failure 404 {object} map[string]string "Task not found"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/tasks/delete/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tasks/delete/"):]
	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)

	err := h.repo.DeleteTask(idNum)

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
