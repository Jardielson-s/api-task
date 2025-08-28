package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	"github.com/Jardielson-s/api-task/modules/shared"
	_ "github.com/Jardielson-s/api-task/modules/tasks/shared"

	"github.com/Jardielson-s/api-task/modules/tasks/model"
	"github.com/Jardielson-s/api-task/modules/tasks/repository"
	"github.com/Jardielson-s/api-task/modules/tasks/services"

	"github.com/go-playground/validator/v10"
)

type CreateTaskBody struct {
	Name    string `json:"name" validate:"required,min=5,max=100" example:"Test 1"`
	Summary string `json:"summary" validate:"required,min=10,max=200" example:"Summary Test 1"`
	Status  string `json:"status" validate:"required,oneof=active inactive pending complete" example:"pending"`
}

type TaskHandler struct {
	service services.TaskService
	repo    repository.TaskRepository
}

var validate = validator.New()

// CreateTask godoc
//
//	@Summary		Create a new task
//	@Description	Create a new task in the database
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security  Bearer
//	@Param			task	body		CreateTaskBody true "Descrição do parâmetro"
//	@Success		201		{object}	shared.CreateTaskResponse  "Task created"
//	@Failure		400		{object}	shared.CreateResponse
//
// @Failure 		400 	{object}	shared.CreateResponse
//
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/tasks [post]
func NewTaskHandler(service services.TaskService, repository repository.TaskRepository) *TaskHandler {
	return &TaskHandler{service, repository}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task CreateTaskBody
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
	tokenInfo, _ := r.Context().Value("tokenInfo").(authenticate.TokenInfo)
	result, err := h.service.CreateTaskService(&model.Task{
		UserId:  tokenInfo.ID,
		Name:    task.Name,
		Summary: task.Summary,
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
