package handler

import (
	"go-ticket-system/internal/handler/user"
	"go-ticket-system/internal/service"
)

// Handler ...
type Handler struct {
	UserHandler user.Handler
}

// New Root Handler
func New(service service.Service) Handler {
	// Init handlers here
	userHandler := user.New(service.UserService)
	return Handler{UserHandler: userHandler}
}
