package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//Server is server
type Server struct {
	httpServer *http.Server
	logger     *zap.SugaredLogger
}

//ApplyHandler is ...
func ApplyHandler(r *mux.Router) http.Handler {
	var httpHandler http.Handler
	httpHandler = handlers.RecoveryHandler()(r)
	return httpHandler
}

// New is ...
func New(r *mux.Router, port int, logger *zap.SugaredLogger) Server {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: ApplyHandler(r),
	}
	return Server{httpServer: httpServer, logger: logger}
}

// Start is ...
func (s Server) Start() error {
	s.logger.Info("Start HTTP Server at " + s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown is ...
func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
