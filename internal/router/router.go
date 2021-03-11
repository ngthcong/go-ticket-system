package router

import (
	"go-ticket-system/internal/handler"
	"go-ticket-system/internal/router/user"
	"go-ticket-system/pkg/common_handler"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Router ...
type Router struct {
	handler handler.Handler
}

// New ...
func New(handler handler.Handler, logger *zap.SugaredLogger) *mux.Router {
	// Create Router
	muxRouter := mux.NewRouter()
	ticketHandler := common_handler.Group(muxRouter, "/api")
	user.New(ticketHandler, handler.UserHandler)

	return muxRouter
}
