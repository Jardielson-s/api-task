package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	"github.com/Jardielson-s/api-task/modules/shared"
)

func findElement(slice []interface{}, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

// ListTasks godoc
//
//	@Summary		List tasks
//	@Description	List tasks in the database
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Security  Bearer
//
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param search query string false "Search query for task name"
// @Success 200 {object} map[string]interface{} "Paginated list of tasks"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/tasks/list [get]
func (h *TaskHandler) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tokenInfo := r.Context().Value("tokenInfo").(authenticate.TokenInfo)
	page := 1
	pageSize := 10
	var search string
	if p := r.URL.Query().Get("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}

	if ps := r.URL.Query().Get("pageSize"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	if sh := r.URL.Query().Get("search"); sh != "" {
		search = sh
	}
	var userId *int
	if findElement(tokenInfo.Roles, shared.GetTechnicianRole()) {
		userId = &tokenInfo.ID
	}
	tasks, totalCount, err := h.repo.ListTasks(page, pageSize, search, userId)
	if err != nil {
		http.Error(w, "Error to load tasks", http.StatusInternalServerError)
		return
	}

	totalPages := (totalCount + int64(pageSize) - 1) / int64(pageSize)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": tasks,
		"meta": map[string]interface{}{
			"totalCount": totalCount,
			"totalPages": totalPages,
			"page":       page,
			"pageSize":   pageSize,
		},
	})
}
