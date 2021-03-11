package ticket

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler ...
type Handler struct {
}

func New() Handler {
	return Handler{}
}

func (th *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusOK, "Hello Quan")
	return
}

// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}
