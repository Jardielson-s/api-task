package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	"github.com/Jardielson-s/api-task/modules/shared"
	"github.com/Jardielson-s/api-task/modules/tasks/repository"
)

// FindTaskById godoc
//
//	@Summary		Find task by id
//	@Description	Find task by id in the database
//	@Security  Bearer
//
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "Task ID"
// @Success 200 {object} shared.CreateTaskResponse "Task data"
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Task not found"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/tasks/{id} [get]
func (h *TaskHandler) FindTaskById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tasks/"):]
	tokenInfo := r.Context().Value("tokenInfo").(authenticate.TokenInfo)

	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)

	var userId *int
	if findElement(tokenInfo.Roles, shared.GetTechnicianRole()) {
		userId = &tokenInfo.ID
	}
	result, err := h.repo.FindByQuery(repository.Query{
		ID:     idNum,
		UserId: userId,
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
	_ = json.NewEncoder(w).Encode(result)
}
