package app

import (
	"context"
	"go-ticket-system/internal/handler"
	"go-ticket-system/internal/repository"
	"go-ticket-system/internal/router"
	"go-ticket-system/internal/service"
	"go-ticket-system/pkg/logger"
	"go-ticket-system/pkg/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"

	"go.uber.org/zap"
)

type (
	app struct {
		server server.Server
		closer io.Closer
		logger *zap.SugaredLogger
	}
	// App interface
	App interface {
		// Start Application
		Start() error
		// Stop Application
		Stop(ctx context.Context) error
		Logger() *zap.SugaredLogger
	}
)

// New Create App instance
func New() App {

	// New log
	logger := logger.New()

	logger.Debug("Start ...")
	defer logger.Debug("End ...")

	// Init DB connection
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Blacklotus@tcp(127.0.0.1:3306)/ticket_system?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connect error, %v", err.Error())
	}
	// Init modules
	repo := repository.New(dbConn, logger)
	s := service.New(repo, logger)
	h := handler.New(s)
	r := router.New(h, logger)

	serv := server.New(r, 8080, logger)

	var c io.Closer

	return app{
		serv,
		c,
		logger,
	}
}

func (app app) Start() error {
	return app.server.Start()
}

func (app app) Stop(ctx context.Context) error {
	if app.closer != nil {
		if err := app.closer.Close(); err != nil {
			return err
		}
	}
	return app.server.Shutdown(ctx)
}

func (app app) Logger() *zap.SugaredLogger {
	return app.logger
}
