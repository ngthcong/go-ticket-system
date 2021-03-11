package user

import (
	"go-ticket-system/internal/handler/user"
	"go-ticket-system/internal/middleware"
	"go-ticket-system/pkg/common_handler"

	"github.com/gorilla/mux"
)

// New ...
func New(muxRouter *mux.Router, userHandler user.Handler) {
	common_handler.POST(muxRouter, "/user", middleware.TokenAuthMiddleware(userHandler.GetUser))
	common_handler.POST(muxRouter, "/login", userHandler.Login)
	common_handler.GET(muxRouter, "/home", middleware.TokenAuthMiddleware(userHandler.UserAsset))
}
