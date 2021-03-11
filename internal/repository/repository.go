package repository

import (
	"go-ticket-system/internal/repository/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	Repository struct {
		UserRepo    user.UserRepository

	}
)

func New(db *gorm.DB, logger *zap.SugaredLogger) Repository {
	userRepo := user.New(db, logger)
	return Repository{
		UserRepo:    userRepo,

	}
}
