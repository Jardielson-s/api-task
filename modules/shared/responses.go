package shared

import (
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"` // Optional field for more specifics
}

func JSONError(w http.ResponseWriter, err CreateResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
