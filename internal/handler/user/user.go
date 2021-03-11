package user

import (
	"encoding/json"
	"go-ticket-system/internal/middleware"
	"go-ticket-system/internal/model"
	"go-ticket-system/internal/service/user"
	"net/http"
)

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler ...
type Handler struct {
	userSev user.UserService
}

func New(userSev user.UserService) Handler {
	return Handler{
		userSev: userSev,
	}
}

func (th *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := middleware.ExtractTokenMetadata(r)
	if err != nil {
		ResponseJSON(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var userRequest model.UserRequest
	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    003,
			ErrorMessage: "Wrong request format",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	user, err := th.userSev.Get(userRequest.Id)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Internal error",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	ResponseJSON(w, http.StatusOK, user)
	return
}

func (th *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenAuth, err := middleware.ExtractTokenMetadata(r)
	if err != nil {
		ResponseJSON(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := th.userSev.Get(tokenAuth.UserId)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Internal error",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	ResponseJSON(w, http.StatusOK, user)
	return
}
func (th *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    003,
			ErrorMessage: "Wrong request format",
		}
		ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	token, err := th.userSev.Login(loginRequest)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Internal error",
		}
		ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	tokens := map[string]string{
		"access_token": token,
	}
	ResponseJSON(w, http.StatusOK, tokens)
	return
}

func (th *Handler) UserAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenAuth, err := middleware.ExtractTokenMetadata(r)

	if err != nil {
		ResponseJSON(w, http.StatusUnauthorized,"")
		return
	}

	assets, err := th.userSev.GetAsset(tokenAuth.UserId)
	if err != nil {
		response := model.ResponseError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Internal error",
		}
		ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}


	ResponseJSON(w, http.StatusOK, assets)
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
