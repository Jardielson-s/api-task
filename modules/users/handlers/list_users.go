package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ListUsers godoc
//
//	@Summary		List users
//	@Description	List users in the database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param search query string false "Search query for username or email"
// @Success 200 {object} map[string]interface{} "Paginated list of users"
//
//	@Failure		500		{string}	string	"Internal Server Error"
//
// @Router			/users/list [get]
func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
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
	users, totalCount, err := h.repo.ListUsers(page, pageSize, search)
	if err != nil {
		http.Error(w, "Error to load users", http.StatusInternalServerError)
		return
	}

	totalPages := (totalCount + int64(pageSize) - 1) / int64(pageSize)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
		"meta": map[string]interface{}{
			"totalCount": totalCount,
			"totalPages": totalPages,
			"page":       page,
			"pageSize":   pageSize,
		},
	})
}
