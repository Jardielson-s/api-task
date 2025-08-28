package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	"github.com/Jardielson-s/api-task/modules/shared"
	"github.com/Jardielson-s/api-task/modules/tasks/model"
)

type UpdateTaskBody struct {
	Name    string  `json:"name" validate:"required,min=4,max=100"`
	Summary *string `json:"summary"`
}

// UpdateTask godoc
//
//	@Summary		Update task
//	@Description	Update task in the database
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security  Bearer
//
// @Param id path int true "Task ID"
//
//	@Param			task	body		UpdateTaskBody true "Description"
//	@Success		201		{object}	shared.CreateTaskResponse  "Task updated"
//	@Failure		400		{object}	shared.CreateResponse
//
// @Failure 		400 	{object}	shared.CreateResponse
//
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/tasks/update/{id} [patch]
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	tokenInfo := r.Context().Value("tokenInfo").(authenticate.TokenInfo)
	var task UpdateTaskBody
	id := r.URL.Path[len("/tasks/update/"):]
	var idNum int
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	fmt.Sscanf(id, "%d", &idNum)
	if err := json.NewDecoder((r.Body)).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(task)
	if err != nil {
		errorResponse := shared.CreateResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: err.Error(),
		}
		shared.JSONError(w, errorResponse, http.StatusBadRequest)
		return
	}
	var userId *int
	if findElement(tokenInfo.Roles, shared.GetTechnicianRole()) {
		userId = &tokenInfo.ID
	}
	result, err := h.service.UpdateTaskService(idNum, model.TaskUpdate{
		Name:    task.Name,
		Summary: task.Summary,
	}, userId)

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
